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

	muxer, err := NewMuxer()
	if err != nil {
		t.Fatalf("Error creating muxer: %v", err)
	}

	if muxer == nil {
		t.Fatal("Muxer is nil")
	}

	if muxer.parser == nil {
		t.Fatal("Parser is nil")
	}

	if muxer.parserV2 == nil {
		t.Fatal("ParserV2 is nil")
	}
}
