package binding

import (
	"fmt"
	"testing"
)

func TestName_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	form := formPostBinding{}
	name := form.Name()
	if name != "form-urlencoded" {
		t.Errorf("Expected 'form-urlencoded', but got '%s'", name)
	}
}
