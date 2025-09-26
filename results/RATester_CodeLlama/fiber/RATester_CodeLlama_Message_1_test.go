package fiber

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	key := "key"
	value := "value"
	level := uint8(1)
	r := &Redirect{
		c: &DefaultCtx{
			flashMessages: redirectionMsgs{
				{
					key:   key,
					value: value,
					level: level,
				},
			},
		},
	}

	// Act
	actual := r.Message(key)

	// Assert
	assert.Equal(t, FlashMessage{
		Key:   key,
		Value: value,
		Level: level,
	}, actual)
}
