package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestJSON_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	// given
	c := &Context{}
	c.JSON(200, "test")

	// when
	c.JSON(200, "test")

	// then
	assert.Equal(t, 200, c.Writer.Status())
}
