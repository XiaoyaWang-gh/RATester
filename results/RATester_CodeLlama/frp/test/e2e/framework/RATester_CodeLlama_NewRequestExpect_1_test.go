package framework

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewRequestExpect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewDefaultFramework()
	reqExpect := NewRequestExpect(f)
	assert.NotNil(t, reqExpect)
}
