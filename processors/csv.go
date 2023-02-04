package common_processors

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/jurelou/forensibus/utils"
	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type CSVProcessor struct {
	processors.Default
	// Input string
}

func (CSVProcessor) Run(in string, _ *processors.Config, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}
	fd, err := os.Open(in)
	if err != nil {
		errors.Add(err)
		return errors
	}
	defer fd.Close()

	csvReader := csv.NewReader(fd)
	csvReader.FieldsPerRecord = -1
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
			errors.Add(err)
			continue
		}
		if len(record) <= 2 {
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
			errors.Add(fmt.Errorf("invalid csv line %v (expected %d columns, got %d)", record, columnsCount, len(record)))
			continue
		}
		row := make(map[string]string, columnsCount)

		for i, col := range record {
			if col != "" {
				row[header[i]] = col
			}
		}
		jsonRow, err := json.Marshal(row)
		if err != nil {
			utils.Log.Warnf("Could not convert csv line to json %v: %s", row, err.Error())
			continue
		}
		e := writer.NewEvent(string(jsonRow))
		out.WriteEvent(e)
	}
	return errors
}

func init() {
	processors.Register("csv", &CSVProcessor{})
}
