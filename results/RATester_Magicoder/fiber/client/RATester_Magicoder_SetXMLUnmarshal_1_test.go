package client

import (
	"fmt"
	"testing"
)

func TestSetXMLUnmarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	xmlUnmarshal := func(data []byte, v any) error {
		// Implement your own XML unmarshal logic here
		return nil
	}

	client.SetXMLUnmarshal(xmlUnmarshal)

	if client.xmlUnmarshal == nil {
		t.Error("xmlUnmarshal is not set")
	}
}
