package aggregator

import (
	"fmt"
	"testing"
)

func TestNewRingChannel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	ch := newRingChannel()

	if ch == nil {
		t.Fatal("expected a non-nil channel")
	}

	if ch.input == nil {
		t.Error("expected a non-nil input channel")
	}

	if ch.output == nil {
		t.Error("expected a non-nil output channel")
	}

	if ch.buffer != nil {
		t.Error("expected a nil buffer")
	}
}
