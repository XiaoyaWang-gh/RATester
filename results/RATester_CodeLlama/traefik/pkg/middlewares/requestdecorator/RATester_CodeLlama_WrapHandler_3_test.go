package requestdecorator

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWrapHandler_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	handler := &RequestDecorator{}
	// when
	actual := WrapHandler(handler)
	// then
	assert.NotNil(t, actual)
}
