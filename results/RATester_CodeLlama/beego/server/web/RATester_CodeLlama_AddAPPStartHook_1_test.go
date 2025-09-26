package web

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddAPPStartHook_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	hf := []hookfunc{func() error {
		return nil
	}}
	// when
	AddAPPStartHook(hf...)
	// then
	assert.Equal(t, hf, hooks)
}
