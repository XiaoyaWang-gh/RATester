package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWeight0_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := pageWithWeight0{weight0: 1}
	assert.Equal(t, 1, p.Weight0())
}
