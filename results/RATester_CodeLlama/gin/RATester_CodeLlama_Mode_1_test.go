package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestMode_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	modeName.Store("test")
	// when
	actual := Mode()
	// then
	assert.Equal(t, "test", actual)
}
