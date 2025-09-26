package postpub

import (
	"fmt"
	"testing"
)

func TestRelPermalink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{
		prefix: "prefix",
	}
	if got := r.RelPermalink(); got != "prefix" {
		t.Errorf("RelPermalink() = %v, want %v", got, "prefix")
	}
}
