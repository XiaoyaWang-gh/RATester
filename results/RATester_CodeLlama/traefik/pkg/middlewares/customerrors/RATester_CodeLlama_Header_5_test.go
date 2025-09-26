package customerrors

import (
	"fmt"
	"testing"
)

func TestHeader_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &codeModifier{}
	if got := r.Header(); got != nil {
		t.Errorf("Header() = %v, want nil", got)
	}
}
