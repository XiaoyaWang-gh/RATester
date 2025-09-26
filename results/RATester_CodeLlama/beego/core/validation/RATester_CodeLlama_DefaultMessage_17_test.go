package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_17(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	a := AlphaNumeric{Key: "AlphaNumeric"}
	if a.DefaultMessage() != MessageTmpls["AlphaNumeric"] {
		t.Errorf("DefaultMessage() = %v, want %v", a.DefaultMessage(), MessageTmpls["AlphaNumeric"])
	}
}
