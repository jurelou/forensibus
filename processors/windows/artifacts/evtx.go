package windows_artifacts

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Velocidex/ordereddict"
	"www.velocidex.com/golang/evtx"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type EvtxProcessor struct {
	processors.Default
}

func (EvtxProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running EVTX processor against %s", in)

	fd, err := os.Open(in)
	if err != nil {
		errors.Add(err)
		return errors
	}
	defer func() {
		if err := fd.Close(); err != nil {
			utils.Log.Warnf("Error while closing file %s: %w", fd.Name(), err)
		}
	}()

	out.SetDefaultSourceType("evtxdump")

	return parseEvtx(fd, out)
}

func parseEvtx(fd *os.File, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	chunks, err := evtx.GetChunks(fd)
	if err != nil {
		errors.Add(err)
		return errors
	}

	for _, chunk := range chunks {
		records, err := chunk.Parse(0)
		if err != nil {
			errors.Add(fmt.Errorf("error while parsing evtx chunk: %w", err))
			continue
		}
		for _, record := range records {

			event := writer.NewEvent("")
			event_map, ok := record.Event.(*ordereddict.Dict)
			if !ok {
				errors.Add(fmt.Errorf("error while reading evtx event: %w", err))
				continue
			}

			eventDict, ok := ordereddict.GetMap(event_map, "Event")
			if !ok {
				errors.Add(fmt.Errorf("error while retrieving evtx Event field"))
				continue
			}
			system, ok := ordereddict.GetMap(eventDict, "System")
			if !ok {
				errors.Add(fmt.Errorf("error while retrieving evtx System field"))
				continue
			}

			userData, ok := ordereddict.GetMap(eventDict, "UserData")
			if ok {
				system.MergeFrom(userData)
			}
			eventData, ok := ordereddict.GetMap(eventDict, "EventData")
			if ok {
				system.MergeFrom(eventData)
			}

			systemEventId, ok := ordereddict.GetMap(system, "EventID")
			if ok {
				eventId, ok := systemEventId.GetInt64("Value")
				if ok {
					system.Set("EventID", eventId)
				}
			}

			systemTimeCreated, ok := ordereddict.GetMap(system, "TimeCreated")
			if ok {
				systemTime, ok := systemTimeCreated.GetInt64("SystemTime")
				if ok {
					event.Time = strconv.Itoa(int(systemTime))
					// system.Set("SystemTime", systemTime)
				}
			}

			eventJson, err := system.MarshalJSON()
			if err != nil {
				errors.Add(fmt.Errorf("could not convert event to json: %w", err))
				continue
			}
			event.Event = string(eventJson)
			out.WriteEvent(event)
		}
	}
	return errors
}

func init() {
	processors.Register("evtxdump", &EvtxProcessor{})
}
