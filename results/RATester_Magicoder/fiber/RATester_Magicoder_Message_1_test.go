package fiber

import (
	"fmt"
	"testing"
)

func TestMessage_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		c: &DefaultCtx{
			flashMessages: redirectionMsgs{
				{
					key:        "test",
					value:      "test",
					level:      1,
					isOldInput: false,
				},
			},
		},
	}

	msg := r.Message("test")

	if msg.Key != "test" || msg.Value != "test" || msg.Level != 1 {
		t.Errorf("Expected message to be 'test', 'test', 1 but got %s, %s, %d", msg.Key, msg.Value, msg.Level)
	}
}
