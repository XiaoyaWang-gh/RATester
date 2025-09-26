package limiter

import (
	"fmt"
	"testing"

	"github.com/tinylib/msgp/msgp"
)

func TestMsgsize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	item := item{
		currHits: 10,
		prevHits: 20,
		exp:      30,
	}

	expectedSize := 1 + 9 + msgp.IntSize + 9 + msgp.IntSize + 4 + msgp.Uint64Size
	actualSize := item.Msgsize()

	if expectedSize != actualSize {
		t.Errorf("Expected size %d, but got %d", expectedSize, actualSize)
	}
}
