package client

import (
	"fmt"
	"testing"
)

func TestSetJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testData := testStruct{
		Name: "John Doe",
		Age:  30,
	}

	req := &Request{}
	req.SetJSON(testData)

	if req.body != testData {
		t.Errorf("Expected body to be %v, got %v", testData, req.body)
	}

	if req.bodyType != jsonBody {
		t.Errorf("Expected bodyType to be %v, got %v", jsonBody, req.bodyType)
	}
}
