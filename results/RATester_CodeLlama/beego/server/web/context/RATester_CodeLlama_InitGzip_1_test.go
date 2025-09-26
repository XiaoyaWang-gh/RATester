package context

import (
	"fmt"
	"testing"
)

func TestInitGzip_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	InitGzip(100, 6, []string{"GET"})
	if gzipMinLength != 100 {
		t.Errorf("gzipMinLength = %d, want %d", gzipMinLength, 100)
	}
	if gzipCompressLevel != 6 {
		t.Errorf("gzipCompressLevel = %d, want %d", gzipCompressLevel, 6)
	}
	if len(includedMethods) != 1 || !includedMethods["GET"] {
		t.Errorf("includedMethods = %v, want %v", includedMethods, []string{"GET"})
	}
}
