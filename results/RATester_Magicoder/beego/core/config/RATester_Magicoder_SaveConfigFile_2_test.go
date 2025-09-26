package config

import (
	"fmt"
	"os"
	"testing"
)

func TestSaveConfigFile_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	iniConfig := &IniConfigContainer{
		data: map[string]map[string]string{
			"section1": {
				"key1": "value1",
				"key2": "value2",
			},
			"section2": {
				"key3": "value3",
				"key4": "value4",
			},
		},
		sectionComment: map[string]string{
			"section1": "This is section1",
			"section2": "This is section2",
		},
		keyComment: map[string]string{
			"section1.key1": "This is key1 in section1",
			"section1.key2": "This is key2 in section1",
			"section2.key3": "This is key3 in section2",
			"section2.key4": "This is key4 in section2",
		},
	}

	err := iniConfig.SaveConfigFile("test.ini")
	if err != nil {
		t.Errorf("SaveConfigFile failed: %v", err)
	}

	// Check if the file exists
	_, err = os.Stat("test.ini")
	if os.IsNotExist(err) {
		t.Errorf("SaveConfigFile failed: file not found")
	}

	// Clean up
	err = os.Remove("test.ini")
	if err != nil {
		t.Errorf("Clean up failed: %v", err)
	}
}
