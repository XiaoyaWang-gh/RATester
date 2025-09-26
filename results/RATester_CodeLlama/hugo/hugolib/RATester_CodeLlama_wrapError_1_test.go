package hugolib

import (
	"errors"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWrapError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	err := p.wrapError(errors.New("test"))
	assert.Equal(t, "test", err.Error())
}
