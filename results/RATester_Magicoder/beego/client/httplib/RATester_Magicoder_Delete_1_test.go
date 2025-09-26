package httplib

import (
	"fmt"
	"testing"
)

func TestDelete_1(t *testing.T) {
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

	var testStruct TestStruct
	testStruct.Field1 = "test"
	testStruct.Field2 = 1

	err := client.Delete(&testStruct, "/test", func(request *BeegoHTTPRequest) {
		request.Header("Content-Type", "application/json")
	})

	if err != nil {
		t.Errorf("Delete failed: %v", err)
	}

	if testStruct.Field1 != "test" || testStruct.Field2 != 1 {
		t.Errorf("Delete failed: expected %v, got %v", testStruct, testStruct)
	}
}
