package httplib

import (
	"fmt"
	"testing"
)

func TestGet_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		Name:     "test",
		Endpoint: "http://localhost:8080",
	}

	var result struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	err := client.Get(&result, "/api/v1/test")
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}

	if result.ID != 1 || result.Name != "test" {
		t.Errorf("Get result is incorrect: %v", result)
	}
}
