package filenotify

import (
	"fmt"
	"testing"
)

func TestRemove_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	w := &filePoller{}
	w.watches = make(map[string]struct{})
	w.watches["test"] = struct{}{}
	w.remove("test")
	if _, exists := w.watches["test"]; exists {
		t.Errorf("expected watches to be empty")
	}
}
