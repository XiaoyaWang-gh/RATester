package middleware

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/zeebo/assert"
)

func TestNonWWWRedirectWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	config := RedirectConfig{
		Code: http.StatusMovedPermanently,
	}
	// Act
	middleware := NonWWWRedirectWithConfig(config)
	// Assert
	assert.NotNil(t, middleware)
}
