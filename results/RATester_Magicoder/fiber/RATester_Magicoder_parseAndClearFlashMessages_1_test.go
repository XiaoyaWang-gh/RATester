package fiber

import (
	"fmt"
	"testing"
)

func TestparseAndClearFlashMessages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		c: &DefaultCtx{
			flashMessages: redirectionMsgs{},
		},
	}

	r.parseAndClearFlashMessages()

	// Add assertions here
}
