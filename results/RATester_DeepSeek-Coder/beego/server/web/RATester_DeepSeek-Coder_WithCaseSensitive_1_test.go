package web

import (
	"fmt"
	"testing"
)

func TestWithCaseSensitive_1(t *testing.T) {
	t.Run("TestWithCaseSensitive", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		opts := &filterOpts{}
		WithCaseSensitive(true)(opts)
		if !opts.routerCaseSensitive {
			t.Errorf("Expected routerCaseSensitive to be true, got false")
		}
		WithCaseSensitive(false)(opts)
		if opts.routerCaseSensitive {
			t.Errorf("Expected routerCaseSensitive to be false, got true")
		}
	})
}
