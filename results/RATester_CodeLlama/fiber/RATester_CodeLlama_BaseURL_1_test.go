package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestBaseURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &DefaultCtx{
		baseURI: "http://example.com",
	}
	assert.Equal(t, c.baseURI, c.BaseURL())
}
