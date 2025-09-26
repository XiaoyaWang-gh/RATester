package postpub

import (
	"fmt"
	"testing"
)

func TestParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{
		prefix: "test",
	}

	params := r.Params()

	if params == nil {
		t.Error("Expected non-nil Params, got nil")
	}
}
