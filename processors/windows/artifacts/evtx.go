package windows_artifacts

import (
	"fmt"
	"os"

	"github.com/Velocidex/ordereddict"
	"www.velocidex.com/golang/evtx"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type EvtxProcessor struct {
	// Input string
}

func (proc EvtxProcessor) Configure() error {
	return nil
}

// func (proc EvtxProcessor) parseEvtx(in string) (PrefetchEntry, error) {
// 	fd, err := os.Open(in)
// 	if err != nil {
// 		utils.Log.Warnf("Could not open prefetch file `%s`", in)
// 		return PrefetchEntry{}, err
// 	}
// 	pf, err := prefetch.LoadPrefetch(fd)
// 	if err != nil {
// 		utils.Log.Warnf("Prefetch file `%s` is invalid: `%s`", in, err.Error())
// 		return PrefetchEntry{}, err
// 	}
// 	return PrefetchEntry(*pf), nil
// }

func (proc EvtxProcessor) Run(in string, out writer.OutputWriter) error {
	fd, err := os.Open(in)
	if err != nil {
		return err
	}
	defer fd.Close()

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

			event_map, ok := record.Event.(*ordereddict.Dict)
			if !ok {
				utils.Log.Warnf("Could not read event")
				continue
			}

			// err := json.Unmarshal([]byte(record.Event.(string)), &event); if err != nil {
			// 	utils.Log.Warnf("Error while converting event to json")
			// 	continue
			// }
			// event_map, ok := record.Event.(*Dict)
			event, ok := ordereddict.GetMap(event_map, "Event")
			if !ok {
				continue
			}
			system, ok := ordereddict.GetMap(event, "System")
			if !ok {
				continue
			}

			userData, ok := ordereddict.GetMap(event, "UserData")
			if ok {
				system.MergeFrom(userData)
			}
			eventData, ok := ordereddict.GetMap(event, "EventData")
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
					system.Set("SystemTime", systemTime)
				}
			}
			json, err := system.MarshalJSON()
			if err != nil {
				return err
			}
			fmt.Println(string(json))

		}
	}
	// utils.Log.Debugf("Run pf processor against `%s`", in)
	// entry, err := proc.parseEvtx(in)
	// if err != nil {
	// 	return err
	// }

	// json, err := json.Marshal(entry)
	// if err != nil {
	// 	return err
	// }

	// out.SetDefaultSourceType("prefetch")
	// e := writer.NewEvent(string(json))
	// out.WriteEvent(e)

	return nil
}

func init() {
	processors.Register("evtxdump", EvtxProcessor{})
}
