package client

import (
	"fmt"
	"os"
	"testing"

	"github.com/fatedier/golib/crypto"
)

func TestInit_12(t *testing.T) {
	t.Run("Testing if DefaultSalt is set correctly", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "frp"
		actual := crypto.DefaultSalt
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("Testing if QUIC_GO_DISABLE_RECEIVE_BUFFER_WARNING is set correctly", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "true"
		actual := os.Getenv("QUIC_GO_DISABLE_RECEIVE_BUFFER_WARNING")
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("Testing if QUIC_GO_DISABLE_ECN is set correctly", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		expected := "true"
		actual := os.Getenv("QUIC_GO_DISABLE_ECN")
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
}
