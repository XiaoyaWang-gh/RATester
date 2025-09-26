package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_11(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	z := ZipCode{}
	if z.DefaultMessage() != MessageTmpls["ZipCode"] {
		t.Errorf("DefaultMessage() = %v, want %v", z.DefaultMessage(), MessageTmpls["ZipCode"])
	}
}
