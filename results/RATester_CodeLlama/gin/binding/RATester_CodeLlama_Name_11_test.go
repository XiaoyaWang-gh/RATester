package binding

import (
	"fmt"
	"testing"
)

func TestName_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	headerBinding := headerBinding{}
	if headerBinding.Name() != "header" {
		t.Errorf("headerBinding.Name() = %v, want %v", headerBinding.Name(), "header")
	}
}
