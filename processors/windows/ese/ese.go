package windows_ese

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Velocidex/ordereddict"
	"www.velocidex.com/golang/go-ese/parser"

	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type ESEProcessor struct {
	processors.Default
}

func (ESEProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	fd, err := os.Open(in)
	if err != nil {
		errors.Add(err)
		return errors
	}
	defer fd.Close()

	ese_ctx, err := parser.NewESEContext(fd)
	if err != nil {
		errors.Add(fmt.Errorf("unable to parse ESE database %s: %w", in, err))
		return errors
	}
	catalog, err := parser.ReadCatalog(ese_ctx)
	if err != nil {
		errors.Add(fmt.Errorf("unable to parse ESE catalog %s: %w", in, err))
		return errors
	}
	for _, name := range catalog.Tables.Keys() {
		catalog.DumpTable(name, func(row *ordereddict.Dict) error {
			jsonRow, err := json.Marshal(row)
			if err != nil {
				return nil
			}
			row.Set("TableName", name)
			e := writer.NewEvent(string(jsonRow))
			out.WriteEvent(e)
			return nil
		})
	}
	return errors
}

func init() {
	processors.Register("ese", ESEProcessor{})
}
