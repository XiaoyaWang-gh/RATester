package binding

import (
	"fmt"
	"testing"
)

func TestName_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formPostBinding := formPostBinding{}
	if formPostBinding.Name() != "form-urlencoded" {
		t.Errorf("formPostBinding.Name() = %v, want %v", formPostBinding.Name(), "form-urlencoded")
	}
}
