package ip

import (
	"fmt"
	"testing"
)

func TestNewChecker_1(t *testing.T) {
	t.Run("should return error when no trusted IPs provided", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := NewChecker([]string{})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error when parsing CIDR trusted IPs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := NewChecker([]string{"invalidCIDR"})
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return Checker with authorized IPs", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		checker, err := NewChecker([]string{"192.0.2.0", "192.0.2.1/32"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if len(checker.authorizedIPs) != 1 {
			t.Errorf("expected 1 authorized IP, got %d", len(checker.authorizedIPs))
		}
		if len(checker.authorizedIPsNet) != 1 {
			t.Errorf("expected 1 authorized IPNet, got %d", len(checker.authorizedIPsNet))
		}
	})
}
