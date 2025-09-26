package httplib

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestJSONMarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testObj := testStruct{
		Name: "John Doe",
		Age:  30,
	}

	b := &BeegoHTTPRequest{}
	jsonBytes, err := b.JSONMarshal(testObj)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	var result testStruct
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}

	if result.Name != testObj.Name || result.Age != testObj.Age {
		t.Errorf("Expected: %v, got: %v", testObj, result)
	}
}
