package idempotency

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMsgsize_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msg := &response{
		Headers: map[string][]string{
			"Content-Type":  {"application/json"},
			"Authorization": {"Bearer token"},
		},
		Body:       []byte("Hello, world!"),
		StatusCode: 200,
	}

	expectedSize := 1 + 3 + msgp.MapHeaderSize +
		2 + msgp.StringPrefixSize + len("Content-Type") + msgp.ArrayHeaderSize + msgp.StringPrefixSize + len("application/json") +
		2 + msgp.StringPrefixSize + len("Authorization") + msgp.ArrayHeaderSize + msgp.StringPrefixSize + len("Bearer token") +
		2 + msgp.BytesPrefixSize + len(msg.Body) +
		3 + msgp.IntSize

	size := msg.Msgsize()
	if size != expectedSize {
		t.Errorf("Expected size %d, but got %d", expectedSize, size)
	}
}
