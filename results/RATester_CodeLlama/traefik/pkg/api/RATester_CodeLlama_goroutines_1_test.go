package api

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoroutines_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	if goroutines() != runtime.NumGoroutine() {
		t.Errorf("goroutines() = %v, want %v", goroutines(), runtime.NumGoroutine())
	}
}
