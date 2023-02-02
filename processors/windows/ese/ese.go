package windows_ese

import (
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
		// table_any, _ := catalog.Tables.Get(name)
		fmt.Println(">", name)
		catalog.DumpTable(name, func(row *ordereddict.Dict) error {
			fmt.Println(" ==", row)
			return nil
		})
		// table := table_any.(*parser.Table)
		// fmt.Println("> ", table.Name)

		// for _, column := range table.Columns {
		// 	fmt.Println("  >", column.Name, column.Type)

		// }

	}
	return errors
}

func init() {
	processors.Register("ese", ESEProcessor{})
}
