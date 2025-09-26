package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveConfigFile_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &IniConfigContainer{
		data: map[string]map[string]string{
			"default": {
				"key1": "val1",
				"key2": "val2",
			},
			"section1": {
				"key1": "val1",
				"key2": "val2",
			},
		},
	}
	filename := "test.ini"
	err := c.SaveConfigFile(filename)
	if err != nil {
		t.Errorf("SaveConfigFile error: %v", err)
	}
	defer os.Remove(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("ReadFile error: %v", err)
	}
	if string(content) != "[default]\nkey1 = val1\nkey2 = val2\n\n[section1]\nkey1 = val1\nkey2 = val2\n" {
		t.Errorf("SaveConfigFile error: %v", string(content))
	}
}
