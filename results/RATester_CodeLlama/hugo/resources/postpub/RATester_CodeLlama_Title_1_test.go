package postpub

import (
	"fmt"
	"testing"
)

func TestTitle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{}
	if got := r.Title(); got != "" {
		t.Errorf("Title() = %v, want %v", got, "")
	}
}
