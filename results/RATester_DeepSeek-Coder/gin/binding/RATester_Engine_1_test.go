package binding

import (
	"fmt"
	"testing"
)

func TestEngine_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	v := &defaultValidator{}
	engine := v.Engine()

	if engine == nil {
		t.Errorf("Expected engine to be not nil, got nil")
	}
}
