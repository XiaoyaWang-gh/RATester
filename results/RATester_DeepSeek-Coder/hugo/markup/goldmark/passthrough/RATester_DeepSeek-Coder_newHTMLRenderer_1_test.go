package passthrough

import (
	"fmt"
	"testing"
)

func TestNewHTMLRenderer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := newHTMLRenderer()
	if r == nil {
		t.Errorf("newHTMLRenderer() = %v, want not nil", r)
	}
}
