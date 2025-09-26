package client

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"
)

func TestAcquireFile_1(t *testing.T) {
	t.Run("Test AcquireFile with setter", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		setter := []SetFileFunc{
			func(f *File) {
				f.SetName("test")
				f.SetPath("/tmp/test")
				f.SetFieldName("file")
				f.SetReader(ioutil.NopCloser(strings.NewReader("test")))
			},
		}
		f := AcquireFile(setter...)
		if f.name != "test" {
			t.Errorf("Expected name to be 'test', got %s", f.name)
		}
		if f.path != "/tmp/test" {
			t.Errorf("Expected path to be '/tmp/test', got %s", f.path)
		}
		if f.fieldName != "file" {
			t.Errorf("Expected fieldName to be 'file', got %s", f.fieldName)
		}
		if f.reader == nil {
			t.Errorf("Expected reader to be not nil")
		}
	})
}
