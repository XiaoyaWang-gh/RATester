package framework

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestExplain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := NewRequestExpect(nil)
	e.Explain("test explain")
	assert.Equal(t, []interface{}{"test explain"}, e.explain)
}
