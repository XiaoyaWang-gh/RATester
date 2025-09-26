package gin

import (
	"fmt"
	"testing"

	"github.com/go-playground/assert"
)

func TestResetColor_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &LogFormatterParams{}
	assert.Equal(t, reset, p.ResetColor())
}
