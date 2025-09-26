package xml

import (
	"fmt"
	"os"
	"testing"
)

func TestSaveConfigFile_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cc := &ConfigContainer{
		data: map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
	}

	err := cc.SaveConfigFile("test.xml")
	if err != nil {
		t.Errorf("SaveConfigFile failed: %v", err)
	}

	_, err = os.Stat("test.xml")
	if os.IsNotExist(err) {
		t.Errorf("SaveConfigFile did not create the file")
	}

	os.Remove("test.xml")
}
