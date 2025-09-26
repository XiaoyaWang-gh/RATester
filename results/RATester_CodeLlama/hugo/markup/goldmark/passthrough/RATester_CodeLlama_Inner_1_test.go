package passthrough

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestInner_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &passthroughContext{}
	p.typ = "inner"
	p.inner = "inner"
	assert.Equal(t, "inner", p.Inner())
}
