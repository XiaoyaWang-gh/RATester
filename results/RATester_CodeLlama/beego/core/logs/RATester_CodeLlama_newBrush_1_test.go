package logs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBrush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	color := "red"
	// when
	brush := newBrush(color)
	// then
	assert.Equal(t, "\033[31m", brush(""))
}
