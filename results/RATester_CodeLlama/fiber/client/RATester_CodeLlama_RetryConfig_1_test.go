package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRetryConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	assert.Equal(t, c.RetryConfig(), nil)
}
