package http

import (
	"fmt"
	"testing"
)

func TestNewMuxer_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	muxer, err := NewMuxer()
	if err != nil {
		t.Fatalf("Error creating muxer: %v", err)
	}

	if muxer.parser == nil {
		t.Error("Parser is nil")
	}

	if muxer.parserV2 == nil {
		t.Error("ParserV2 is nil")
	}

	if muxer.defaultHandler == nil {
		t.Error("DefaultHandler is nil")
	}
}
