package ese

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Velocidex/ordereddict"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/writer"
	"github.com/jurelou/forensibus/utils/processors"
)

var tablesGUIDs = map[string]string{
	"{5C8CF1C7-7257-4F13-B223-970EF5939312}": "Execution",
	"{973F5D5C-1D90-4944-BE8E-24B94231A174}": "Network usage",
	"{DD6636C4-8929-4683-974E-22C046A43763}": "Network connections",
	"{D10CA2FE-6FCF-4F6D-848E-B2E99266FA89}": "Application resource usage",
}

type SrumProcessor struct {
	processors.Default
}

func (SrumProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	utils.Log.Infof("Running SRUM processor against %s", in)

	fd, err := os.Open(in)
	if err != nil {
		errors.Add(fmt.Errorf("could not open ese file %s: %w", in, err))
		return errors
	}
	defer func() {
		if err := fd.Close(); err != nil {
			utils.Log.Warnf("Error while closing file %s: %w", fd.Name(), err)
		}
	}()

	catalog, err := GetCatalog(fd)
	if err != nil {
		errors.Add(err)
		return errors
	}

	idMap := GetIdMap(catalog)

	for guid, tableName := range tablesGUIDs {
		catalog.DumpTable(guid, func(row *ordereddict.Dict) error {
			appId, exists := row.GetInt64("AppId")
			if exists {
				resolvedAppId, idExists := idMap[appId]
				if idExists {
					row.Set("AppId", resolvedAppId)
				}
			}
			userId, exists := row.GetInt64("UserId")
			if exists {
				resolvedUserId, idExists := idMap[userId]
				if idExists {
					row.Set("UserId", resolvedUserId)
				}
			}
			row.Set("TableName", tableName)

			jsonRow, err := json.Marshal(row)
			if err != nil {
				return nil
			}
			e := writer.NewEvent(string(jsonRow))
			timeAny, timeExists := row.Get("TimeStamp")
			if timeExists {
				switch timeStamp := timeAny.(type) {
				case time.Time:
					e.Time = strconv.FormatInt(timeStamp.Unix(), 10)
				case string:
					e.Time = timeStamp
				}
			}
			out.WriteEvent(e)
			return nil
		})
	}
	return errors
}

func init() {
	processors.Register("srum", &SrumProcessor{})
}
