package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPage_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := pageWithWeight0{}
	assert.Equal(t, p.page(), p.pageState)
}
