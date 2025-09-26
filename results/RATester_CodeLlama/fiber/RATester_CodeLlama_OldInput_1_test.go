package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestOldInput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	key := "key"
	value := "value"
	r := &Redirect{
		c: &DefaultCtx{
			flashMessages: redirectionMsgs{
				{
					key:        key,
					value:      value,
					isOldInput: true,
				},
			},
		},
	}

	// When
	oldInput := r.OldInput(key)

	// Then
	assert.Equal(t, key, oldInput.Key)
	assert.Equal(t, value, oldInput.Value)
}
