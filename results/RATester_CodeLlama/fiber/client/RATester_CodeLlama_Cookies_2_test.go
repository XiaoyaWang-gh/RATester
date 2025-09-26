package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestCookies_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Response{}
	// when
	cookies := r.Cookies()
	// then
	assert.Nil(t, cookies)
}
