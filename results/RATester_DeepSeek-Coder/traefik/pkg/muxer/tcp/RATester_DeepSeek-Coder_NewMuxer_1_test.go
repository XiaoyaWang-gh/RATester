package tcp

import (
	"fmt"
	"testing"
)

func TestNewMuxer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	muxer, err := NewMuxer()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if muxer == nil {
		t.Fatalf("Expected muxer not to be nil")
	}

	if muxer.parser == nil {
		t.Errorf("Expected parser not to be nil")
	}

	if muxer.parserV2 == nil {
		t.Errorf("Expected parserV2 not to be nil")
	}
}
