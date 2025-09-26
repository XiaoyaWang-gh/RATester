package postpub

import (
	"fmt"
	"testing"
)

func TestName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{}
	if got := r.Name(); got != "" {
		t.Errorf("Name() = %v, want %v", got, "")
	}
}
