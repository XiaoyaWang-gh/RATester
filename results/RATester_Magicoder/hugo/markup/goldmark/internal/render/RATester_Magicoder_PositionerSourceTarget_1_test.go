package render

import (
	"fmt"
	"testing"
)

func TestPositionerSourceTarget_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hookBase{
		getSourceSample: func() []byte {
			return []byte("test")
		},
	}

	result := hb.PositionerSourceTarget()

	if string(result) != "test" {
		t.Errorf("Expected 'test', got '%s'", string(result))
	}
}
