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

	lf := &localFile{
		NativePath: "testdata/test.txt",
		fs:         afero.NewOsFs(),
	}
	if got := lf.MD5(); !bytes.Equal(got, []byte("12345678901234567890123456789012")) {
		t.Errorf("MD5() = %v, want %v", got, []byte("12345678901234567890123456789012"))
	}
}
