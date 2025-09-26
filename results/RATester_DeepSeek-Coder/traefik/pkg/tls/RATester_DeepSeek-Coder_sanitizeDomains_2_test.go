package tls

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestSanitizeDomains_2(t *testing.T) {
	t.Run("should return error when no domain was given", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := sanitizeDomains(types.Domain{})
		if err == nil {
			t.Errorf("Expected error, got nil")
		}
	})

	t.Run("should return clean domains", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		domains := types.Domain{
			Main: "example.com",
			SANs: []string{"www.example.com", "blog.example.com"},
		}
		expected := []string{"example.com", "www.example.com", "blog.example.com"}
		cleanDomains, err := sanitizeDomains(domains)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !reflect.DeepEqual(cleanDomains, expected) {
			t.Errorf("Expected %v, got %v", expected, cleanDomains)
		}
	})
}
