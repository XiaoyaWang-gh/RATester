package binding

import (
	"fmt"
	"testing"
)

func TestName_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formMultipartBinding := formMultipartBinding{}
	if formMultipartBinding.Name() != "multipart/form-data" {
		t.Errorf("formMultipartBinding.Name() = %v, want %v", formMultipartBinding.Name(), "multipart/form-data")
	}
}
