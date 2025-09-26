package net

import (
	"fmt"
	"testing"
)

func TestMakeHTTPGzipHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HTTPGzipWrapper{}
	handler := MakeHTTPGzipHandler(h)
	if handler == nil {
		t.Errorf("MakeHTTPGzipHandler() = %v, want %v", handler, h)
	}
}
