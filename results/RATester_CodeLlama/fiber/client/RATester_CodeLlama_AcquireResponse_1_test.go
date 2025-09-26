package client

import (
	"fmt"
	"sync"
	"testing"

	"github.com/zeebo/assert"
)

func TestAcquireResponse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	responsePool := sync.Pool{
		New: func() interface{} {
			return &Response{}
		},
	}
	responsePool.Put(&Response{})
	// Act
	resp := AcquireResponse()
	// Assert
	assert.NotNil(t, resp)
}
