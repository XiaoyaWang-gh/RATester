package commands

import (
	"fmt"
	"net/url"
	"strings"
	"testing"
)

func TestInjectLiveReloadScript_1(t *testing.T) {
	t.Run("test with a valid reader and baseURL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		src := strings.NewReader("<html><body></body></html>")
		baseURL, _ := url.Parse("http://localhost:35729")
		expected := "<html><body><script src=\"http://localhost:35729/livereload.js?snipver=1\"></script></body></html>"
		result := injectLiveReloadScript(src, baseURL)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("test with an invalid reader", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		src := strings.NewReader("")
		baseURL, _ := url.Parse("http://localhost:35729")
		expected := ""
		result := injectLiveReloadScript(src, baseURL)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})

	t.Run("test with an invalid baseURL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		src := strings.NewReader("<html><body></body></html>")
		baseURL, _ := url.Parse("")
		expected := "<html><body></body></html>"
		result := injectLiveReloadScript(src, baseURL)
		if result != expected {
			t.Errorf("Expected %s, got %s", expected, result)
		}
	})
}
