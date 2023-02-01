package windows_ese

import (
	"os"
	"fmt"

	// "www.velocidex.com/golang/go-ese/parser"
	"github.com/Velocidex/ordereddict"

	"github.com/jurelou/forensibus/utils/processors"
	"github.com/jurelou/forensibus/utils/writer"
)

type SrumProcessor struct {
	processors.Default
}

func (SrumProcessor) Run(in string, out writer.OutputWriter) processors.PError {
	errors := processors.PError{}

	fd, err := os.Open(in)
	if err != nil {
		errors.Add(fmt.Errorf("could not open ese file %s: %w", in, err))
		return errors
	}
	defer fd.Close()

	catalog, err := GetCatalog(fd)
	if err != nil {
		errors.Add(err)
		return errors
	}
	// var id_map map[int64]string

	err = catalog.DumpTable("SruDbIdMapTable", func(row *ordereddict.Dict) error {
		idType, exists := row.GetInt64("IdType")
		if !exists {
			return nil
		}
		idIndex, exists := row.GetInt64("IdIndex")
		if !exists {
			return nil
		}
		idBlob, exists := row.GetString("IdBlob")
		if !exists {
			return nil
		}
		fmt.Println(">>>>>>>", idType, idIndex, idBlob)
		if idType == 3 {
			fmt.Println("====", FormatGUID(idBlob))
		} else {

		}
		// id_details := &SRUMId{}
		// err := arg_parser.ExtractArgsWithContext(ctx, scope, row, id_details)
		// if err != nil {
		// 	return err
		// }

		// // Its a GUID
		// if id_details.IdType == 3 {
		// 	id_details.IdBlob = formatGUID(id_details.IdBlob)
		// } else {
		// 	id_details.IdBlob = formatString(id_details.IdBlob)
		// }

		// lookup_map[id_details.IdIndex] = id_details.IdBlob
		return nil
	})
	// for _, name := range catalog.Tables.Keys() {
	// 	// table_any, _ := catalog.Tables.Get(name)
	// 	fmt.Println(">", name)
	// 	catalog.DumpTable(name, func(row *ordereddict.Dict) error {
	// 		fmt.Println(" ==", row)
	// 		return nil
	// 	})
	// 	// table := table_any.(*parser.Table)
	// 	// fmt.Println("> ", table.Name)

	// 	// for _, column := range table.Columns {
	// 	// 	fmt.Println("  >", column.Name, column.Type)

	// 	// }

	// }
	return errors
}

func init() {
	processors.Register("srum", SrumProcessor{})
}
