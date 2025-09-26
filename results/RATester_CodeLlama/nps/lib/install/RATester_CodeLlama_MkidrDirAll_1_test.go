package install

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestMkidrDirAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "/tmp/test"
	v := []string{"a", "b", "c"}
	MkidrDirAll(path, v...)
	if _, err := os.Stat(path); err != nil {
		t.Errorf("Failed to create directory %s error:%s", path, err.Error())
	}
	if _, err := os.Stat(filepath.Join(path, "a")); err != nil {
		t.Errorf("Failed to create directory %s error:%s", filepath.Join(path, "a"), err.Error())
	}
	if _, err := os.Stat(filepath.Join(path, "b")); err != nil {
		t.Errorf("Failed to create directory %s error:%s", filepath.Join(path, "b"), err.Error())
	}
	if _, err := os.Stat(filepath.Join(path, "c")); err != nil {
		t.Errorf("Failed to create directory %s error:%s", filepath.Join(path, "c"), err.Error())
	}
}
