package postpub

import (
	"fmt"
	"testing"
)

func TestResourceType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{}
	if got := r.ResourceType(); got != "" {
		t.Errorf("ResourceType() = %v, want %v", got, "")
	}
}
