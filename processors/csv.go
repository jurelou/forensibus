package common_processors

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"io"
	"fmt"
	// "time"


	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)


type CSVProcessor struct {
	// Input string
}

func (proc CSVProcessor) Configure() error {
	return nil
}

// func (proc CSVProcessor) parseCsv(in string) (PrefetchEntry, error) {
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

func (proc CSVProcessor) Run(in string, out writer.OutputWriter) error {
	utils.Log.Debugf("Run csv processor against `%s`", in)
    fd, err := os.Open(in)
    if err != nil {
        return err
    }
    defer fd.Close()

	csvReader := csv.NewReader(fd)
	csvReader.FieldsPerRecord  = -1
	csvReader.Comment = '#'
	csvReader.Comma = ','

	var header []string
	var columnsCount int

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if len(record) <=2 {
			// Skip if line does not contains at least 3 columns.
			// This helps to skip potential unwanted file headers
			continue
		}
		if header == nil {
			// First row, set it as a header
			header = record
			columnsCount = len(header)
			continue
		}
		if len(record) != columnsCount {
			return fmt.Errorf("Invalid csv line %v (expected %d columns, got %d)", record, columnsCount, len(record))
		}
		row := make(map[string]string, columnsCount)

		for i, col := range record{
			if col != "" {
				row[header[i]] = col

			}
		}
		jsonRow, err := json.Marshal(row); if err != nil {
			utils.Log.Warnf("Could not convert csv line to json %v: %s", row, err.Error())
			continue
		}
		e := writer.NewEvent(string(jsonRow))
		out.WriteEvent(e)
	}
	return nil
}

func init() {
	processors.Register("csv", CSVProcessor{})
}
