package fiber

import (
	"fmt"
	"testing"
)

func TestOldInputs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		c: &DefaultCtx{
			flashMessages: redirectionMsgs{
				{
					key:        "testKey",
					value:      "testValue",
					isOldInput: true,
				},
			},
		},
	}

	inputs := r.OldInputs()

	if len(inputs) != 1 {
		t.Errorf("Expected 1 input, got %d", len(inputs))
	}

	if inputs[0].Key != "testKey" {
		t.Errorf("Expected key 'testKey', got '%s'", inputs[0].Key)
	}

	if inputs[0].Value != "testValue" {
		t.Errorf("Expected value 'testValue', got '%s'", inputs[0].Value)
	}
}
