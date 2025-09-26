package fiber

import (
	"fmt"
	"testing"
)

func TestWith_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Redirect{
		messages: redirectionMsgs{
			{key: "test", value: "old", level: 1},
		},
	}

	r.With("test", "new", 2)

	if len(r.messages) != 1 {
		t.Errorf("Expected 1 message, got %d", len(r.messages))
	}

	if r.messages[0].value != "new" {
		t.Errorf("Expected value 'new', got '%s'", r.messages[0].value)
	}

	if r.messages[0].level != 2 {
		t.Errorf("Expected level 2, got %d", r.messages[0].level)
	}
}
