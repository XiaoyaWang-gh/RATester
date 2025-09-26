package middleware

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRewrite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	rules := map[string]string{
		"^/api/v1/users/([0-9]+)$": "/api/v1/users/$1",
	}
	// when
	middleware := Rewrite(rules)
	// then
	assert.NotNil(t, middleware)
}
