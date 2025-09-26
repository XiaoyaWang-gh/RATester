package tplimpl

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAddTemplate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var tplHandler templateHandler
	var name string
	var tpl string

	// when
	err := tplHandler.AddTemplate(name, tpl)

	// then
	assert.NoError(t, err)
}
