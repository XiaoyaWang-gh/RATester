package framework

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestNewFramework_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	opt := Options{
		TotalParallelNode: 1,
		CurrentNodeIndex:  0,
		FromPortIndex:     10000,
		ToPortIndex:       10000,
	}
	f := NewFramework(opt)
	assert.NotNil(t, f)
}
