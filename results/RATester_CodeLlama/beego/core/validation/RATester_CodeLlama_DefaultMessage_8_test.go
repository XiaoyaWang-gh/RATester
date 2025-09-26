package validation

import (
	"fmt"
	"testing"
)

func TestDefaultMessage_8(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := Mobile{}
	want := MessageTmpls["Mobile"]
	got := m.DefaultMessage()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
