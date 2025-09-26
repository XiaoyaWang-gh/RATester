package yaml

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestSaveConfigFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &ConfigContainer{
		data: map[string]interface{}{
			"a": "1",
			"b": "2",
		},
	}
	filename := "test.yaml"
	err := c.SaveConfigFile(filename)
	if err != nil {
		t.Errorf("SaveConfigFile error %v", err)
	}
	defer os.Remove(filename)
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Errorf("ReadFile error %v", err)
	}
	if string(content) != "a: 1\nb: 2\n" {
		t.Errorf("SaveConfigFile error %v", string(content))
	}
}
