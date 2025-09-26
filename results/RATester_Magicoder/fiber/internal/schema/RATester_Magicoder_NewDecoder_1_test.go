package schema

import (
	"fmt"
	"testing"
)

func TestNewDecoder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	decoder := NewDecoder()
	if decoder == nil {
		t.Error("Expected a non-nil decoder, but got nil")
	}
	if decoder.cache == nil {
		t.Error("Expected a non-nil cache, but got nil")
	}
	if decoder.zeroEmpty != false {
		t.Error("Expected zeroEmpty to be false, but got true")
	}
	if decoder.ignoreUnknownKeys != false {
		t.Error("Expected ignoreUnknownKeys to be false, but got true")
	}
}
