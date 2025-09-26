package gin

import (
	"fmt"
	"testing"
)

func TestSetMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msg := &Error{}
	msg.SetMeta(1)
	if msg.Meta != 1 {
		t.Errorf("SetMeta failed")
	}
}
