package middleware

import (
	"fmt"
	"net/http"
	"testing"
)

func TestWriteHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &ignorableWriter{
		ignoreWrites: false,
	}

	w.WriteHeader(http.StatusOK)

	if w.ignoreWrites != false {
		t.Errorf("Expected ignoreWrites to be false, but got %v", w.ignoreWrites)
	}
}
