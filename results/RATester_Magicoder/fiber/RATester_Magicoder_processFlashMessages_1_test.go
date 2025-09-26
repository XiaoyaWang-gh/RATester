package fiber

import (
	"fmt"
	"testing"
)

func TestprocessFlashMessages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		messages: redirectionMsgs{},
	}

	r.processFlashMessages()

	if len(r.messages) != 0 {
		t.Errorf("Expected messages to be empty, but got %d", len(r.messages))
	}
}
