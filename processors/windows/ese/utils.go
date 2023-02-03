package ese

import (
	"fmt"
	"os"

	"github.com/Velocidex/ordereddict"
	"www.velocidex.com/golang/go-ese/parser"
)

func GetCatalog(fd *os.File) (*parser.Catalog, error) {
	ese_ctx, err := parser.NewESEContext(fd)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ESE database %s: %w", fd.Name(), err)
	}
	catalog, err := parser.ReadCatalog(ese_ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to parse ESE catalog %s: %w", fd.Name(), err)
	}
	return catalog, nil
}

func GetIdMap(catalog *parser.Catalog) map[int64]string {
	idMap := make(map[int64]string)

	catalog.DumpTable("SruDbIdMapTable", func(row *ordereddict.Dict) error {
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
		if idType == 3 {
			idMap[idIndex] = FormatSID(idBlob)
		} else {
			idMap[idIndex] = FormatUtf16String(idBlob)
		}
		return nil
	})

	return idMap
}
