package binding

import (
	"fmt"
	"testing"
)

func TestName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	jsonBinding := jsonBinding{}
	if jsonBinding.Name() != "json" {
		t.Errorf("jsonBinding.Name() = %v, want %v", jsonBinding.Name(), "json")
	}
}
