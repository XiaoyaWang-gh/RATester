package httplib

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestToXML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type TestStruct struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	testStruct := TestStruct{
		Name: "John",
		Age:  30,
	}

	xmlData, err := xml.Marshal(testStruct)
	if err != nil {
		t.Fatalf("Failed to marshal test struct to XML: %v", err)
	}

	req := &BeegoHTTPRequest{
		body: xmlData,
	}

	var result TestStruct
	err = req.ToXML(&result)
	if err != nil {
		t.Fatalf("Failed to unmarshal XML to test struct: %v", err)
	}

	if result.Name != testStruct.Name || result.Age != testStruct.Age {
		t.Fatalf("Unmarshalled struct does not match original struct")
	}
}
