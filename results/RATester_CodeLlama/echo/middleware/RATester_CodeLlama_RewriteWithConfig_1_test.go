package middleware

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestRewriteWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	config := RewriteConfig{
		Rules: map[string]string{
			"/old":              "/new",
			"/api/*":            "/$1",
			"/js/*":             "/public/javascripts/$1",
			"/users/*/orders/*": "/user/$1/order/$2",
		},
	}
	assert.NotPanics(t, func() {
		RewriteWithConfig(config)
	})
}
