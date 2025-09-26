package middleware

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestHTTPSWWWRedirectWithConfig_1(t *testing.T) {
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
	actual := HTTPSWWWRedirectWithConfig(config)
	// THEN
	assert.NotNil(t, actual)
}
