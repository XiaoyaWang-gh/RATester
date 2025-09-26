package binding

import (
	"fmt"
	"testing"
)

func TestName_13(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	queryBinding := queryBinding{}
	if queryBinding.Name() != "query" {
		t.Errorf("queryBinding.Name() = %v, want %v", queryBinding.Name(), "query")
	}
}
