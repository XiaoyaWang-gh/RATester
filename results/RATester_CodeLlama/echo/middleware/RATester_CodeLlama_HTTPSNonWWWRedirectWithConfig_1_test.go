package middleware

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestHTTPSNonWWWRedirectWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// GIVEN
	config := RedirectConfig{
		Code: http.StatusMovedPermanently,
	}
	// WHEN
	middleware := HTTPSNonWWWRedirectWithConfig(config)
	// THEN
	assert.NotNil(t, middleware)
}
