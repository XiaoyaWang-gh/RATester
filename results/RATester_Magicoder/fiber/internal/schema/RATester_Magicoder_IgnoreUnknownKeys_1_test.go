package schema

import (
	"fmt"
	"testing"
)

func TestIgnoreUnknownKeys_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Decoder{}
	d.IgnoreUnknownKeys(true)
	if d.ignoreUnknownKeys != true {
		t.Errorf("Expected ignoreUnknownKeys to be true, but got %v", d.ignoreUnknownKeys)
	}
}
