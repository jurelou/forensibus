package decompress

import (
	_ "fmt"
	"testing"
)

// func TestInvalidFile(t *testing.T) {
// 	err := DecompressSevenZip("./this_does_not_exists.xz", "/tmp/out")
// 	if err == nil {
// 		t.Errorf("Decompress should fail on invalid file: `./this_does_not_exists.xz`")
// 	}

// }

// func TestEncrypted(t *testing.T) {
// 	file := "../../datasets/archives/EmptyEncryptedPasswdIsPasswd.7z"
// 	err := DecompressSevenZip(file, "/tmp/out")
// 	if err != nil {
// 		t.Errorf("Decompress archive `EmptyEncryptedPasswdIsPasswd.7z` should not fail")
// 	}
// 	enc, err := IsEncrypted(file); if err != nil {
// 		t.Errorf("Failed while checking if file `EmptyEncryptedPasswdIsPasswd` is encrypted: %s", err.Error())
// 	}
// 	if enc == false {
// 		t.Errorf("`EmptyEncryptedPasswdIsPasswd` is encrypted ...")
// 	}

// }

func TestTotoa(t *testing.T) {
	DecompressSevenZip("/tmp/ok2/EXE_TMP.7z", "/tmp/ok")
	// DecompressSevenZip("/tmp/ok2/EXE_TMP.7z", "/tmp/ok3")
}
