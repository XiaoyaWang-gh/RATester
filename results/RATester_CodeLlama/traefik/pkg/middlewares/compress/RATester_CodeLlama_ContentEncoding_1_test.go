package compress

import (
	"fmt"
	"testing"
)

func TestContentEncoding_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &compressionWriter{alg: "gzip"}
	if c.ContentEncoding() != "gzip" {
		t.Errorf("ContentEncoding() = %v, want %v", c.ContentEncoding(), "gzip")
	}
}
