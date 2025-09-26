package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestPureJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// given
	c := &Context{}
	c.PureJSON(200, "test")

	// when
	c.PureJSON(200, "test")

	// then
	assert.Equal(t, 200, c.Writer.Status())
}
