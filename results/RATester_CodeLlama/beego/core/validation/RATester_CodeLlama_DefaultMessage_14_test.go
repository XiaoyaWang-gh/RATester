package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := Base64{}
	want := MessageTmpls["Base64"]
	got := b.DefaultMessage()
	if got != want {
		t.Errorf("DefaultMessage() = %v, want %v", got, want)
	}
}
