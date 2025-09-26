package compare

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestConditional_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	n := &Namespace{}
	cond := true
	v1 := "v1"
	v2 := "v2"
	assert.Equal(t, v1, n.Conditional(cond, v1, v2))
	cond = false
	assert.Equal(t, v2, n.Conditional(cond, v1, v2))
}
