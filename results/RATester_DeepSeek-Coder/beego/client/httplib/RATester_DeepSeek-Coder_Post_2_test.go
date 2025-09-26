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

	type testStruct struct {
		Field1 string `json:"field1"`
		Field2 int    `json:"field2"`
	}

	testBody := testStruct{
		Field1: "test",
		Field2: 1,
	}

	var result testStruct
	err := client.Post(&result, "/test", testBody)
	if err != nil {
		t.Errorf("Post() error = %v", err)
		return
	}

	if result.Field1 != "test" {
		t.Errorf("Post() result.Field1 = %v, want %v", result.Field1, "test")
	}

	if result.Field2 != 1 {
		t.Errorf("Post() result.Field2 = %v, want %v", result.Field2, 1)
	}
}
