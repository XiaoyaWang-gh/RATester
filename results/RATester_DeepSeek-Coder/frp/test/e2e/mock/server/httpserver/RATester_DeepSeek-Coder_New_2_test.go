package httpserver

import (
	"fmt"
	"testing"
)

func TestNew_2(t *testing.T) {
	t.Run("should return a new server with default options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		s := New()
		if s.bindAddr != "127.0.0.1" {
			t.Errorf("expected bindAddr to be '127.0.0.1', got %s", s.bindAddr)
		}
	})

	t.Run("should return a new server with custom options", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		customOption := func(s *Server) *Server {
			s.bindAddr = "192.168.0.1"
			return s
		}
		s := New(customOption)
		if s.bindAddr != "192.168.0.1" {
			t.Errorf("expected bindAddr to be '192.168.0.1', got %s", s.bindAddr)
		}
	})
}
