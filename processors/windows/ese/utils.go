package windows_ese

import (
	"fmt"
	"os"

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
