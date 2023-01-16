package windows_artifacts

import (
	"os"
	"strconv"

	"github.com/Velocidex/ordereddict"
	"www.velocidex.com/golang/evtx"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type EvtxProcessor struct{}

func (proc EvtxProcessor) Configure() error {
	return nil
}

func (proc EvtxProcessor) Run(in string, out writer.OutputWriter) error {
	fd, err := os.Open(in)
	if err != nil {
		return err
	}
	defer fd.Close()

	out.SetDefaultSourceType("evtxdump")

	return parseEvtx(fd, out)
}

func parseEvtx(fd *os.File, out writer.OutputWriter) error {
	chunks, err := evtx.GetChunks(fd)
	if err != nil {
		return err
	}
	for _, chunk := range chunks {
		records, err := chunk.Parse(0)
		if err != nil {
			utils.Log.Warnf("Error while parsing chunk")
			continue
		}
		for _, record := range records {

			event := writer.NewEvent("")

			event_map, ok := record.Event.(*ordereddict.Dict)
			if !ok {
				utils.Log.Warnf("Could not read event")
				continue
			}

			eventDict, ok := ordereddict.GetMap(event_map, "Event")
			if !ok {
				return nil
			}
			system, ok := ordereddict.GetMap(eventDict, "System")
			if !ok {
				return nil
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
				utils.Log.Warnf("Could not convert event to json")
				continue
			}
			event.Event = string(eventJson)
			out.WriteEvent(event)
		}
	}
	return nil
}

func init() {
	processors.Register("evtxdump", EvtxProcessor{})
}
