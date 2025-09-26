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
		t.Errorf("Expected NewDecoder to return a non-nil value")
	}
	if decoder.cache == nil {
		t.Errorf("Expected Decoder.cache to be non-nil")
	}
	if decoder.zeroEmpty {
		t.Errorf("Expected Decoder.zeroEmpty to be false by default")
	}
	if decoder.ignoreUnknownKeys {
		t.Errorf("Expected Decoder.ignoreUnknownKeys to be false by default")
	}
}
