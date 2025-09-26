package hugolib

import (
	"fmt"
	"testing"
)

func TestTestOptWithConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		fn func(c *IntegrationTestConfig)
		c  *IntegrationTestConfig
	)

	fn = func(c *IntegrationTestConfig) {
		// do something with c
	}

	opt := TestOptWithConfig(fn)
	opt(c)
}
