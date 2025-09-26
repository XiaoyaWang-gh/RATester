package client

import (
	"fmt"
	"testing"
)

func TestSetJSONUnmarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	testFunc := func(data []byte, v any) error {
		return nil
	}

	client.SetJSONUnmarshal(testFunc)

	if client.jsonUnmarshal == nil {
		t.Error("jsonUnmarshal is nil")
	}
}
