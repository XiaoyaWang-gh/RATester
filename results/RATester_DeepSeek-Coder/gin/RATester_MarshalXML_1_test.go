package gin

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"testing"
)

func TestMarshalXML_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := H{
		"key1": "value1",
		"key2": "value2",
	}

	expected := `<map><key1>value1</key1><key2>value2</key2></map>`
	b := &bytes.Buffer{}
	enc := xml.NewEncoder(b)
	err := h.MarshalXML(enc, xml.StartElement{})
	if err != nil {
		t.Errorf("MarshalXML() error = %v", err)
		return
	}
	if got := b.String(); got != expected {
		t.Errorf("MarshalXML() = %v, want %v", got, expected)
	}
}
