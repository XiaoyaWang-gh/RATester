package plugins

import (
	zipa "archive/zip"
	"fmt"
	"testing"
)

func TestUnzipFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &zipa.File{}
	dest := ""
	err := unzipFile(f, dest)
	if err != nil {
		t.Errorf("unzipFile() error = %v, wantErr %v", err, nil)
		return
	}
}
