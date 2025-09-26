package admin

import (
	"fmt"
	"testing"
)

func TestIsSuccess_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	r := &Result{Status: 200}
	if !r.IsSuccess() {
		t.Errorf("IsSuccess() = %v, want %v", r.IsSuccess(), true)
	}
}
