package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWidth_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := &errorResource{}
	assert.Panics(t, func() {
		e.Width()
	})
}
