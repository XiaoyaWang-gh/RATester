package session

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMsgsize_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &data{
		Data: map[string]any{
			"key1": "value1",
			"key2": 2,
			"key3": true,
		},
	}

	expectedSize := 1 + 5 + msgp.MapHeaderSize +
		msgp.StringPrefixSize + len("key1") + msgp.GuessSize("value1") +
		msgp.StringPrefixSize + len("key2") + msgp.GuessSize(2) +
		msgp.StringPrefixSize + len("key3") + msgp.GuessSize(true)

	size := d.Msgsize()

	if size != expectedSize {
		t.Errorf("Expected size %d, but got %d", expectedSize, size)
	}
}
