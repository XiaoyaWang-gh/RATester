package cache

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMsgsize_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := &item{
		headers: map[string][]byte{
			"header1": []byte("value1"),
			"header2": []byte("value2"),
		},
		body:      []byte("body"),
		ctype:     []byte("application/json"),
		cencoding: []byte("utf-8"),
		status:    200,
		exp:       1630496341,
	}

	expectedSize := 1 + 8 + msgp.MapHeaderSize +
		(msgp.StringPrefixSize + len("header1") + msgp.BytesPrefixSize + len("value1") +
			msgp.StringPrefixSize + len("header2") + msgp.BytesPrefixSize + len("value2")) +
		msgp.BytesPrefixSize + len("body") +
		msgp.BytesPrefixSize + len("application/json") +
		msgp.BytesPrefixSize + len("utf-8") +
		7 + msgp.IntSize +
		4 + msgp.Uint64Size +
		8 + msgp.IntSize

	size := item.Msgsize()

	if size != expectedSize {
		t.Errorf("Expected size %d, but got %d", expectedSize, size)
	}
}
