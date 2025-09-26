package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestInitQueryCache_1(t *testing.T) {
	t.Run("Test initQueryCache with nil Request and URL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		context := &Context{
			Request: nil,
		}
		context.initQueryCache()
		if context.queryCache != nil {
			t.Errorf("Expected queryCache to be nil, got %v", context.queryCache)
		}
	})

	t.Run("Test initQueryCache with non-nil Request and nil URL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		context := &Context{
			Request: &http.Request{
				URL: nil,
			},
		}
		context.initQueryCache()
		if context.queryCache != nil {
			t.Errorf("Expected queryCache to be nil, got %v", context.queryCache)
		}
	})

	t.Run("Test initQueryCache with non-nil Request and URL", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		context := &Context{
			Request: &http.Request{
				URL: &url.URL{
					RawQuery: "test=1",
				},
			},
		}
		context.initQueryCache()
		if context.queryCache == nil {
			t.Errorf("Expected queryCache to be non-nil")
		}
		if context.queryCache.Get("test") != "1" {
			t.Errorf("Expected queryCache to contain 'test=1', got %v", context.queryCache)
		}
	})
}
