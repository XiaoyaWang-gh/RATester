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

	req := &Request{}
	testXML := struct {
		Name string `xml:"name"`
	}{"Test"}

	req.SetXML(testXML)

	if req.body != testXML {
		t.Errorf("Expected body to be %v, but got %v", testXML, req.body)
	}

	if req.bodyType != xmlBody {
		t.Errorf("Expected bodyType to be %v, but got %v", xmlBody, req.bodyType)
	}
}
