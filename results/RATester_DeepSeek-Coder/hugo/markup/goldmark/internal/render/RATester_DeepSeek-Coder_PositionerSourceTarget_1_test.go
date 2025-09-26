package render

import (
	"bytes"
	"fmt"
	"testing"
)

func TestPositionerSourceTarget_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &hookBase{
		getSourceSample: func() []byte {
			return []byte("test")
		},
	}

	result := h.PositionerSourceTarget()
	expected := []byte("test")

	if !bytes.Equal(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
