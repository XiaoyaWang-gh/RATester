package page

import (
	"fmt"
	"testing"
)

func TestNewLazyContentProvider_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := func() (OutputFormatContentProvider, error) {
		return nil, nil
	}
	lcp := NewLazyContentProvider(f)
	if lcp == nil {
		t.Error("NewLazyContentProvider returned nil")
	}
}
