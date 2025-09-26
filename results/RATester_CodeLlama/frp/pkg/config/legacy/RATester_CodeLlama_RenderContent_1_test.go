package legacy

import (
	"fmt"
	"testing"
)

func TestRenderContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := []byte(`{{.Name}}`)
	out, err := RenderContent(in)
	if err != nil {
		t.Errorf("RenderContent error: %v", err)
		return
	}
	if string(out) != "test" {
		t.Errorf("RenderContent error: %v", string(out))
	}
}
