package gin

import (
	"fmt"
	"testing"
)

func TestDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := Dir("./", true)
	if fs == nil {
		t.Error("Dir() should not return nil")
	}
}
