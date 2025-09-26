package transport

import (
	"fmt"
	"testing"
)

func TestNewCertPool_1(t *testing.T) {
	t.Run("valid_ca_path", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		pool, err := newCertPool("valid_ca_path")
		if err != nil {
			t.Errorf("newCertPool() error = %v, wantErr nil", err)
			return
		}
		if pool == nil {
			t.Errorf("newCertPool() = nil, want non-nil")
		}
	})

	t.Run("invalid_ca_path", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		pool, err := newCertPool("invalid_ca_path")
		if err == nil {
			t.Errorf("newCertPool() error = nil, wantErr not nil")
		}
		if pool != nil {
			t.Errorf("newCertPool() = non-nil, want nil")
		}
	})
}
