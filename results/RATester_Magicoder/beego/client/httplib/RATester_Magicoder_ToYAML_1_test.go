package httplib

import (
	"fmt"
	"testing"

	"gopkg.in/yaml.v3"
)

func TestToYAML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Field1 string `yaml:"field1"`
		Field2 int    `yaml:"field2"`
	}

	testObj := testStruct{
		Field1: "test",
		Field2: 123,
	}

	yamlBytes, err := yaml.Marshal(testObj)
	if err != nil {
		t.Fatalf("Failed to marshal test object to YAML: %v", err)
	}

	req := &BeegoHTTPRequest{
		body: yamlBytes,
	}

	var result testStruct
	err = req.ToYAML(&result)
	if err != nil {
		t.Fatalf("Failed to unmarshal YAML to test object: %v", err)
	}

	if result.Field1 != testObj.Field1 || result.Field2 != testObj.Field2 {
		t.Errorf("Unmarshalled object does not match original object. Expected: %v, Got: %v", testObj, result)
	}
}
