package echo

import (
	"fmt"
	"testing"
)

func TestUpdateAllowHeader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &routeMethods{}
	m.updateAllowHeader()
	if m.allowHeader != "" {
		t.Errorf("Expected empty string, got %s", m.allowHeader)
	}
}
