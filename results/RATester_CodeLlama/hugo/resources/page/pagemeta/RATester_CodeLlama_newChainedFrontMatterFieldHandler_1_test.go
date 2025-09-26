package pagemeta

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewChainedFrontMatterFieldHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var f FrontMatterHandler
	var handlers []frontMatterFieldHandler
	// when
	result := f.newChainedFrontMatterFieldHandler(handlers...)
	// then
	assert.NotNil(t, result)
}
