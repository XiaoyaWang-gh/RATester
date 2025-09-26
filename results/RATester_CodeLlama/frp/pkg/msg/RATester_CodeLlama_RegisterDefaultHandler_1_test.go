package msg

import (
	"fmt"
	"testing"
)

func TestRegisterDefaultHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := NewDispatcher(nil)
	d.RegisterDefaultHandler(func(m Message) {})
	if d.defaultHandler == nil {
		t.Fatal("defaultHandler should not be nil")
	}
}
