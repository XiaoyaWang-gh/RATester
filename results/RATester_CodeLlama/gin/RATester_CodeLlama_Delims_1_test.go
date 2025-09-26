package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestDelims_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	engine := New()
	engine.Delims("{{", "}}")
	assert.Equal(t, "{{", engine.delims.Left)
	assert.Equal(t, "}}", engine.delims.Right)
}
