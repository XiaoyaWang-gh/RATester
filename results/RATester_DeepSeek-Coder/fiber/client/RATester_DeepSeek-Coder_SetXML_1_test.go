package client

import (
	"fmt"
	"testing"
)

func TestSetXML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	type testStruct struct {
		Name string `xml:"name"`
		Age  int    `xml:"age"`
	}

	testData := testStruct{
		Name: "John Doe",
		Age:  30,
	}

	req := &Request{}
	req.SetXML(testData)

	if req.body != testData {
		t.Errorf("Expected body to be %v, got %v", testData, req.body)
	}

	if req.bodyType != xmlBody {
		t.Errorf("Expected bodyType to be %v, got %v", xmlBody, req.bodyType)
	}
}
