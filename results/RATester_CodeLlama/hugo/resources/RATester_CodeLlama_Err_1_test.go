package resources

import (
	"fmt"
	"testing"
)

func TestErr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &genericResource{}
	if got := r.Err(); got != nil {
		t.Errorf("Err() = %v, want nil", got)
	}
}
