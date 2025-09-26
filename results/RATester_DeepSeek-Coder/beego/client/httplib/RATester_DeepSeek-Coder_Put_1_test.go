package httplib

import (
	"fmt"
	"testing"
)

func TestPut_1(t *testing.T) {
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

	var result testStruct
	err := client.Put(&result, "/test", testStruct{Field1: "test", Field2: 123})
	if err != nil {
		t.Errorf("Put() error = %v", err)
		return
	}

	if result.Field1 != "test" {
		t.Errorf("Put() result.Field1 = %v, want %v", result.Field1, "test")
	}

	if result.Field2 != 123 {
		t.Errorf("Put() result.Field2 = %v, want %v", result.Field2, 123)
	}
}
