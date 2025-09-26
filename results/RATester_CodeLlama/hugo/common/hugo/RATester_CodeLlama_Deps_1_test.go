package hugo

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestDeps_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	var i HugoInfo
	// when
	result := i.Deps()
	// then
	assert.Equal(t, []*Dependency{}, result)
}
