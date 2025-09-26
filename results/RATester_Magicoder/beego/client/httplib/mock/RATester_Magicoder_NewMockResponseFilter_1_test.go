package mock

import (
	"fmt"
	"testing"
)

func TestNewMockResponseFilter_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	filter := NewMockResponseFilter()
	if filter == nil {
		t.Error("NewMockResponseFilter() should not return nil")
	}
	if len(filter.ms) != 0 {
		t.Error("NewMockResponseFilter() should initialize ms with an empty slice")
	}
}
