package echo

import (
	"fmt"
	"testing"
)

func TestUnwrap_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Response{}
	w := r.Unwrap()
	if w != r.Writer {
		t.Errorf("Unwrap() = %v, want %v", w, r.Writer)
	}
}
