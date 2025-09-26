package gin

import (
	"errors"
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestCustomRecovery_1(t *testing.T) {
	t.Run("TestCustomRecovery", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		mockCtx := &Context{}
		mockErr := errors.New("mock error")
		mockRecoveryFunc := func(c *Context, err any) {
			assert.Equal(t, mockCtx, c)
			assert.Equal(t, mockErr, err)
		}

		handler := CustomRecovery(mockRecoveryFunc)
		handler(mockCtx)

		assert.Equal(t, mockCtx, mockCtx)
		assert.Equal(t, mockErr, mockErr)
	})
}
