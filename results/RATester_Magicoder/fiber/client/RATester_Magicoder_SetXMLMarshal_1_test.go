package client

import (
	"fmt"
	"testing"
)

func TestSetXMLMarshal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{}

	testXMLMarshal := func(v any) ([]byte, error) {
		return []byte("<test>test</test>"), nil
	}

	client.SetXMLMarshal(testXMLMarshal)

	if client.xmlMarshal == nil {
		t.Error("XMLMarshal not set")
	}

	xml, err := client.xmlMarshal(struct{}{})
	if err != nil {
		t.Error(err)
	}

	if string(xml) != "<test>test</test>" {
		t.Error("XMLMarshal not working as expected")
	}
}
