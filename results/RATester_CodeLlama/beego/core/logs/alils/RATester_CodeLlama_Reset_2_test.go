package alils

import (
	"fmt"
	"testing"
)

func TestReset_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &LogGroup{}
	m.Reset()
	if m.Logs != nil {
		t.Errorf("Logs should be nil")
	}
	if m.Reserved != nil {
		t.Errorf("Reserved should be nil")
	}
	if m.Source != nil {
		t.Errorf("Source should be nil")
	}
	if m.Topic != nil {
		t.Errorf("Topic should be nil")
	}
}
