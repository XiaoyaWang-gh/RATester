package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestInitPage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// CONTEXT_0
	p := &pageState{}
	// CONTEXT_1
	p.initPage()
	// ASSERT
	assert.Equal(t, 0, p.pageOutputIdx)
}
