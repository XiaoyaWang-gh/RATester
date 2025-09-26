package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_21(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := IP{}
	want := MessageTmpls["IP"]
	got := i.DefaultMessage()
	if got != want {
		t.Errorf("IP.DefaultMessage() = %v, want %v", got, want)
	}
}
