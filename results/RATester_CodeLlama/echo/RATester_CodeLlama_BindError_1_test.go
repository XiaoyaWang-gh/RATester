package echo

import (
	"errors"
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBindError_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	b := &ValueBinder{}
	b.errors = []error{errors.New("error")}
	// when
	err := b.BindError()
	// then
	assert.Equal(t, errors.New("error"), err)
}
