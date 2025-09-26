package middleware

import (
	"fmt"
	"testing"
)

func TestUnwrap_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &bodyDumpResponseWriter{}
	if w.Unwrap() != w.ResponseWriter {
		t.Errorf("w.Unwrap() = %v, want %v", w.Unwrap(), w.ResponseWriter)
	}
}
