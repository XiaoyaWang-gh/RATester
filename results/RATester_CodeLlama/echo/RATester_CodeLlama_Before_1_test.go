package echo

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestBefore_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Response{}
	fn := func() {}
	// when
	r.Before(fn)
	// then
	assert.Equal(t, []func(){fn}, r.beforeFuncs)
}
