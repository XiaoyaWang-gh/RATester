package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReleaseResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	resp := &Response{}
	// when
	ReleaseResponse(resp)
	// then
	assert.Equal(t, nil, resp.client)
	assert.Equal(t, nil, resp.request)
}
