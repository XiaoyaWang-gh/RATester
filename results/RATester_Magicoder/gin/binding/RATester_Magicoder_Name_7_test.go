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
	name := form.Name()
	if name != "form" {
		t.Errorf("Expected 'form', but got '%s'", name)
	}
}
