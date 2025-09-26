package deploy

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestMD5_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	afero.WriteFile(fs, "/test.txt", []byte("test content"), 0644)
	lf := &localFile{
		NativePath: "/test.txt",
		fs:         fs,
	}

	md5 := lf.MD5()
	expectedMD5 := []byte{0x9e, 0x37, 0x79, 0xb9, 0x7f, 0x4a, 0x7c, 0x15, 0xab, 0x8f, 0xc, 0x4e, 0x96, 0x4e, 0x42, 0x87}
	if !bytes.Equal(md5, expectedMD5) {
		t.Errorf("Expected MD5 hash to be %v, got %v", expectedMD5, md5)
	}
}
