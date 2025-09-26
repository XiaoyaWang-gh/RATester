package conn

import (
	"fmt"
	"testing"
	"time"
)

func TestNewOptions_1(t *testing.T) {
	t.Parallel()

	t.Run("default", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		opts := newOptions()
		if opts.Timeout != defaultTimeOut {
			t.Errorf("expected default timeout to be %v, got %v", defaultTimeOut, opts.Timeout)
		}
	})

	t.Run("custom", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		customTimeout := 5 * time.Second
		opts := newOptions(func(o *Options) {
			o.Timeout = customTimeout
		})
		if opts.Timeout != customTimeout {
			t.Errorf("expected custom timeout to be %v, got %v", customTimeout, opts.Timeout)
		}
	})
}
