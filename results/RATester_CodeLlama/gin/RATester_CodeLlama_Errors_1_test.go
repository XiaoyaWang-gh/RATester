package gin

import (
	"fmt"
	"testing"
)

func TestErrors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var a errorMsgs
	if got := a.Errors(); got != nil {
		t.Errorf("Errors() = %v, want nil", got)
	}
}
