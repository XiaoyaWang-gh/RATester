package template

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPrepare_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	template := &Template{}

	// when
	result, err := template.Prepare()

	// then
	assert.NoError(t, err)
	assert.Equal(t, template, result)
}
