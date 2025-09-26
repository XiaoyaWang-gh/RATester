package fiber

import (
	"fmt"
	"testing"
)

func TestOldInput_1(t *testing.T) {
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

	expected := OldInputData{
		Key:   "testKey",
		Value: "testValue",
	}

	result := r.OldInput("testKey")

	if result != expected {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
