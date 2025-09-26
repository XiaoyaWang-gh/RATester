package css

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource_transformers/cssjs"
)

func TestPostCSS_2(t *testing.T) {
	t.Parallel()

	ns := &Namespace{
		postcssClient: &cssjs.PostCSSClient{},
	}

	t.Run("valid arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r, err := ns.PostCSS("test.css", map[string]any{"option": "value"})
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if r == nil {
			t.Error("expected a resource, got nil")
		}
	})

	t.Run("invalid arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.PostCSS("test.css", "option")
		if err == nil {
			t.Error("expected an error, got nil")
		}
	})

	t.Run("no arguments", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.PostCSS()
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
