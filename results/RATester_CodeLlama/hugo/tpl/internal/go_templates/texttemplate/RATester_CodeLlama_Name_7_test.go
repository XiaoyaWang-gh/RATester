package template

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestName_7(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	template := &Template{
		name: "test",
	}
	// when
	result := template.Name()
	// then
	assert.Equal(t, "test", result)
}
