package middleware

import (
	"fmt"
	"testing"
)

func TestGzipDecompressPool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &DefaultGzipDecompressPool{}
	p := d.gzipDecompressPool()
	if p.New == nil {
		t.Errorf("p.New is nil")
	}
}
