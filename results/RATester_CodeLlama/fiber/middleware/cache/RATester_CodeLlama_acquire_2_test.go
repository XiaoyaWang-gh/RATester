package cache

import (
	"fmt"
	"sync"
	"testing"

	"github.com/zeebo/assert"
)

func TestAcquire_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// setup
	m := &manager{
		pool: sync.Pool{
			New: func() interface{} {
				return &item{}
			},
		},
	}
	// exercise
	actual := m.acquire()
	// verify
	assert.NotNil(t, actual)
}
