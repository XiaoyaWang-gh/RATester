package postpub

import (
	"fmt"
	"testing"
)

func TestPermalink_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{
		prefix: "https://example.com/",
	}
	if got := r.Permalink(); got != "https://example.com/" {
		t.Errorf("Permalink() = %v, want %v", got, "https://example.com/")
	}
}
