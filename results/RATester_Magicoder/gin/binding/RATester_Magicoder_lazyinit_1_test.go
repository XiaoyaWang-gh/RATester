package binding

import (
	"fmt"
	"testing"
)

func Testlazyinit_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	v := &defaultValidator{}
	v.lazyinit()

	if v.validate == nil {
		t.Error("Expected validate to be initialized, but it was nil")
	}
}
