package binding

import (
	"fmt"
	"testing"
)

func TestName_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	form := formBinding{}
	if form.Name() != "form" {
		t.Errorf("form.Name() = %v, want %v", form.Name(), "form")
	}
}
