package postpub

import (
	"context"
	"fmt"
	"testing"
)

func TestContent_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{
		prefix: "prefix",
	}
	got, err := r.Content(context.Background())
	if err != nil {
		t.Fatalf("Content() error = %v", err)
	}
	if got != "prefix" {
		t.Errorf("Content() got = %v, want %v", got, "prefix")
	}
}
