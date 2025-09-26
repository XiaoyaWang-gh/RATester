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
					key:        "test",
					value:      "test",
					isOldInput: true,
				},
			},
		},
	}

	expected := OldInputData{
		Key:   "test",
		Value: "test",
	}

	result := r.OldInput("test")

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
