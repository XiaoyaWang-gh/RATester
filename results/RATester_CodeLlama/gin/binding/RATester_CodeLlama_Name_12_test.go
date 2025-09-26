package binding

import (
	"fmt"
	"testing"
)

func TestName_12(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	uriBinding := uriBinding{}
	if uriBinding.Name() != "uri" {
		t.Errorf("uriBinding.Name() = %v, want %v", uriBinding.Name(), "uri")
	}
}
