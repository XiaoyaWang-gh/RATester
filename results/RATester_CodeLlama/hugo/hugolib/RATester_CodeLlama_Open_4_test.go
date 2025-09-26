package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestOpen_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	fi := &fileInfo{}

	// when
	f, err := fi.Open()

	// then
	assert.Nil(t, f)
	assert.NotNil(t, err)
}
