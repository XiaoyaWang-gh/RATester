package config

import (
	"fmt"
	"testing"
)

func TestSaveConfigFile_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &fakeConfigContainer{}
	filename := "test.conf"
	err := c.SaveConfigFile(filename)
	if err == nil {
		t.Errorf("SaveConfigFile() error = %v, wantErr %v", err, "not implement in the fakeConfigContainer")
	}
}
