package postpub

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &PostPublishResource{}
	assert.PanicsWithValue(t, "field not supported: Params", func() {
		r.Params()
	})
}
