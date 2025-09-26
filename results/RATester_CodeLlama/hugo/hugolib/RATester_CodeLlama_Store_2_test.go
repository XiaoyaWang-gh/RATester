package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestStore_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageCommon{}
	assert.Equal(t, p.store, p.Store())
}
