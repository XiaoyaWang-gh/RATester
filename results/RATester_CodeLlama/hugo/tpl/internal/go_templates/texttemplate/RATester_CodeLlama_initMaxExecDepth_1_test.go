package template

import (
	"fmt"
	"testing"
)

func TestInitMaxExecDepth_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if initMaxExecDepth() != 100000 {
		t.Errorf("initMaxExecDepth() = %v, want %v", initMaxExecDepth(), 100000)
	}
}
