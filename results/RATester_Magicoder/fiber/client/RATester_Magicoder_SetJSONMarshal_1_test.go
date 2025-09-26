package client

import (
	"fmt"
	"testing"
)

func TestSetJSONMarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}
	jsonMarshal := func(v any) ([]byte, error) {
		return []byte("test"), nil
	}
	client.SetJSONMarshal(jsonMarshal)
	if client.jsonMarshal == nil {
		t.Error("jsonMarshal is nil")
	}
}
