package web

import (
	"fmt"
	"os"
	"testing"
)

func TestSaveConfigFile_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &beegoAppConfig{}
	filename := "test.conf"
	err := b.SaveConfigFile(filename)
	if err != nil {
		t.Errorf("SaveConfigFile() error = %v", err)
		return
	}
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		t.Errorf("SaveConfigFile() error = %v", err)
		return
	}
	os.Remove(filename)
}
