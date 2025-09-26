package alils

import (
	"fmt"
	"testing"
)

func TestGetContents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var m *Log
	if m.GetContents() != nil {
		t.Errorf("GetContents() = %v, want nil", m.GetContents())
	}
}
