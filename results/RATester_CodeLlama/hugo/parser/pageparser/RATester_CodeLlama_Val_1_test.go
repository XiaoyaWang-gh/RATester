package pageparser

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestVal_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	source := []byte("abcdefghijklmnopqrstuvwxyz")
	i := Item{
		low:  1,
		high: 5,
	}

	// when
	result := i.Val(source)

	// then
	assert.Equal(t, []byte("bcde"), result)
}
