package middleware

import (
	"fmt"
	"testing"
)

func TestClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &limitedReader{}
	r.Reset(nil)
	if err := r.Close(); err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}
