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
		messages: []redirectionMsg{
			{key: "test1", value: "value1", level: 1, isOldInput: false},
			{key: "test2", value: "value2", level: 2, isOldInput: true},
		},
	}

	r.With("test3", "value3", 3)

	if len(r.messages) != 3 {
		t.Errorf("Expected length of messages to be 3, got %d", len(r.messages))
	}

	expected := redirectionMsg{key: "test3", value: "value3", level: 3, isOldInput: false}
	if r.messages[2] != expected {
		t.Errorf("Expected message to be %v, got %v", expected, r.messages[2])
	}
}
