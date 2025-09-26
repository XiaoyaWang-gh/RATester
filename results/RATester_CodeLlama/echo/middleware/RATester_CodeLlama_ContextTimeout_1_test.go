package middleware

import (
	"fmt"
	"testing"
	"time"

	"github.com/zeebo/assert"
)

func TestContextTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	timeout := time.Duration(10)
	middleware := ContextTimeout(timeout)
	assert.NotNil(t, middleware)
}
