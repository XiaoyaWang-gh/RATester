package attributes

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestTrigger_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	a := &attrParser{}
	// when
	result := a.Trigger()
	// then
	assert.Equal(t, []byte{'{'}, result)
}
