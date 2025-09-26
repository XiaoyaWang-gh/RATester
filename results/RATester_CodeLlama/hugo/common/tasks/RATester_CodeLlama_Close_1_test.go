package tasks

import (
	"fmt"
	"testing"
)

func TestClose_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &RunEvery{}
	if err := r.Close(); err != nil {
		t.Errorf("Close() = %v, want nil", err)
	}
}
