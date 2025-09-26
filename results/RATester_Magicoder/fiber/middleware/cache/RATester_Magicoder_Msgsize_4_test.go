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
		headers:   map[string][]byte{"key": []byte("value")},
		body:      []byte("body"),
		ctype:     []byte("ctype"),
		cencoding: []byte("cencoding"),
		status:    200,
		exp:       100,
	}

	expectedSize := 1 + 8 + msgp.MapHeaderSize + msgp.StringPrefixSize + len("key") + msgp.BytesPrefixSize + len("value") + msgp.BytesPrefixSize + len("body") + msgp.BytesPrefixSize + len("ctype") + msgp.BytesPrefixSize + len("cencoding") + 7 + msgp.IntSize + 4 + msgp.Uint64Size + 8 + msgp.IntSize

	size := item.Msgsize()

	if size != expectedSize {
		t.Errorf("Expected size %d, but got %d", expectedSize, size)
	}
}
