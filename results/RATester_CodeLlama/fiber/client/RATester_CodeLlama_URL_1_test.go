package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestURL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.url = "https://www.baidu.com"
	assert.Equal(t, r.URL(), "https://www.baidu.com")
}
