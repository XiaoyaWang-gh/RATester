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

	req := &Request{}
	testValue := struct {
		Name string `json:"name"`
	}{
		Name: "test",
	}
	req.SetJSON(testValue)

	if req.body != testValue {
		t.Errorf("Expected body to be %v, but got %v", testValue, req.body)
	}

	if req.bodyType != jsonBody {
		t.Errorf("Expected bodyType to be %v, but got %v", jsonBody, req.bodyType)
	}
}
