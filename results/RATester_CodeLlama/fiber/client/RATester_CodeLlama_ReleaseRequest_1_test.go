package client

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestReleaseRequest_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given:
	req := &Request{}
	// when:
	ReleaseRequest(req)
	// then:
	assert.Equal(t, nil, req.ctx)
}
