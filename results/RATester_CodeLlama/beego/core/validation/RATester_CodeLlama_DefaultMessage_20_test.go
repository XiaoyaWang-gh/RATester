package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := Numeric{Key: "test"}
	if n.DefaultMessage() != MessageTmpls["Numeric"] {
		t.Errorf("DefaultMessage() = %v, want %v", n.DefaultMessage(), MessageTmpls["Numeric"])
	}
}
