package css

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource_transformers/cssjs"
)

func TestPostCSS_2(t *testing.T) {
	ns := &Namespace{
		postcssClient: &cssjs.PostCSSClient{},
	}

	t.Run("success", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		r, err := ns.PostCSS(nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if r == nil {
			t.Error("expected resource, got nil")
		}
	})

	t.Run("error", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := ns.PostCSS(nil, nil, nil)
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
