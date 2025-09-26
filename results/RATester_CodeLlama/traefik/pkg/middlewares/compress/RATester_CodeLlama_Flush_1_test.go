package compress

import (
	"fmt"
	"testing"
)

func TestFlush_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var r responseWriter

	r.Flush()

	if r.seenData {
		t.Error("Flush() should not set seenData")
	}

	if r.compressionStarted {
		t.Error("Flush() should not set compressionStarted")
	}

	if r.buf != nil {
		t.Error("Flush() should not set buf")
	}

	if r.statusCodeSet {
		t.Error("Flush() should not set statusCodeSet")
	}

	if r.statusCode != 0 {
		t.Error("Flush() should not set statusCode")
	}
}
