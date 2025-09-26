package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestListenerAddr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	assert.Nil(t, e.ListenerAddr())
}
