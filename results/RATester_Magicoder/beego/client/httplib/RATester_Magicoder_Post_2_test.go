package httplib

import (
	"fmt"
	"testing"
)

func TestPost_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Name:     "test",
		Endpoint: "http://localhost:8080",
	}

	type TestStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testStruct := TestStruct{
		Field1: "test",
		Field2: 1,
	}

	err := client.Post(&testStruct, "/test", testStruct)
	if err != nil {
		t.Errorf("Post failed: %v", err)
	}

	if testStruct.Field1 != "test" || testStruct.Field2 != 1 {
		t.Errorf("Post failed: expected %v, got %v", testStruct, testStruct)
	}
}
