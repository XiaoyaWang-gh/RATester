package json

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestSaveConfigFile_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	container := &JSONConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	tempFile, err := ioutil.TempFile("", "test*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	err = container.SaveConfigFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to save config file: %v", err)
	}

	data, err := ioutil.ReadFile(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	var expectedData map[string]interface{}
	err = json.Unmarshal(data, &expectedData)
	if err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	if !reflect.DeepEqual(container.data, expectedData) {
		t.Errorf("Expected data to be %v, but got %v", container.data, expectedData)
	}
}
