package gin

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoveryWithWriter_1(t *testing.T) {
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		out := &bytes.Buffer{}
		recovery := func(c *Context, err any) {}
		// when
		handler := RecoveryWithWriter(out, recovery)
		// then
		assert.NotNil(t, handler)
	})
	t.Run("test case 2", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		// given
		out := &bytes.Buffer{}
		recovery := func(c *Context, err any) {}
		// when
		handler := RecoveryWithWriter(out, recovery)
		// then
		assert.NotNil(t, handler)
	})
}
