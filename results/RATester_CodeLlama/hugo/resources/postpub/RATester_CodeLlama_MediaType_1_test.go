package postpub

import (
	"fmt"
	"testing"
)

func TestMediaType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{}
	m := r.MediaType()
	if len(m) != 0 {
		t.Errorf("MediaType() = %v, want %v", m, 0)
	}
}
