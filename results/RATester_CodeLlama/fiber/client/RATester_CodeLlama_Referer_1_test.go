package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReferer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Request{}
	r.referer = "https://www.baidu.com"
	assert.Equal(t, "https://www.baidu.com", r.Referer())
}
