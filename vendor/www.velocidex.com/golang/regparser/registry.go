package regparser

import (
	"errors"
	"io"
)

// Model a registry hive with this object.
type Registry struct {
	Reader io.ReaderAt

	Profile   *RegistryProfile
	BaseBlock *HBASE_BLOCK
}

func NewRegistry(reader io.ReaderAt) (*Registry, error) {
	self := &Registry{
		Reader:  reader,
		Profile: NewRegistryProfile(),
	}
	self.BaseBlock = self.Profile.HBASE_BLOCK(reader, 0)
	if self.BaseBlock.Signature() != 0x66676572 {
		return nil, errors.New("File does not have registry magic.")
	}

	return self, nil
}

// A helper method to open a key by path.
func (self *Registry) OpenKey(key_path string) *CM_KEY_NODE {
	root_cell := self.Profile.HCELL(self.Reader,
		0x1000+int64(self.BaseBlock.RootCell()))

	nk := root_cell.KeyNode()
	if nk == nil {
		return nil
	}

subkey_match:
	for _, component := range SplitComponents(key_path) {
		if component == "" {
			continue
		}

		for _, subkey := range nk.Subkeys() {
			if subkey.Name() == component {
				nk = subkey
				continue subkey_match
			}
		}

		// If we get here we could not find the key:
		return nil
	}

	return nk
}
