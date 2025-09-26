package httplib

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWithProxy_1(t *testing.T) {
	t.Parallel()

	t.Run("should return a ClientOption", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		t.Parallel()

		proxy := func(*http.Request) (*url.URL, error) {
			return nil, nil
		}

		result := WithProxy(proxy)

		assert.IsType(t, func(*Client) {}, result)
	})
}
