package config

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestMarshalJSON_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Field1 string
		Field2 int
	}

	testStructInstance := testStruct{
		Field1: "test",
		Field2: 1,
	}

	ns := ConfigNamespace[testStruct, testStruct]{
		SourceStructure: testStructInstance,
		SourceHash:      "hash",
		Config:          testStructInstance,
	}

	expectedJSON, _ := json.Marshal(testStructInstance)

	jsonBytes, err := ns.MarshalJSON()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if string(jsonBytes) != string(expectedJSON) {
		t.Errorf("Expected JSON: %s, but got: %s", string(expectedJSON), string(jsonBytes))
	}
}
