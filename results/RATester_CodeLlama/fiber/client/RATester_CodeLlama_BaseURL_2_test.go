package client

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBaseURL_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Client{}
	c.SetBaseURL("https://example.com")
	assert.Equal(t, "https://example.com", c.BaseURL())
}
