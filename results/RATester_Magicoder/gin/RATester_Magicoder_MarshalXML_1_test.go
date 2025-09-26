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

	h := H{"key1": "value1", "key2": "value2"}
	buf := &bytes.Buffer{}
	enc := xml.NewEncoder(buf)
	err := h.MarshalXML(enc, xml.StartElement{})
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	enc.Flush()
	expected := `<map><key1>value1</key1><key2>value2</key2></map>`
	if buf.String() != expected {
		t.Errorf("Expected: %s, got: %s", expected, buf.String())
	}
}
