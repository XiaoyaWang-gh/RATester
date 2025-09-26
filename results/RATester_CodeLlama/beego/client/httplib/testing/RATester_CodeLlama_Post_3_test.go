package testing

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestPost_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "/test"
	req := Post(path)
	assert.Equal(t, baseURL+getPort()+path, req.GetRequest().URL.String())
}
