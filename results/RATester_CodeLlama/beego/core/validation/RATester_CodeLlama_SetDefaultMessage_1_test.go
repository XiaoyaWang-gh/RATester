package validation

import (
	"fmt"
	"testing"
)

func TestSetDefaultMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msg := map[string]string{
		"test": "test",
	}
	SetDefaultMessage(msg)
	if len(MessageTmpls) != 1 {
		t.Errorf("SetDefaultMessage failed")
	}
}
