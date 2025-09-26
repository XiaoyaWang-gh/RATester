package try

import (
	"fmt"
	"os"
	"testing"
	"time"
)

func TestApplyCIMultiplier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	timeout := time.Duration(10)
	ci := "true"
	os.Setenv("CI", ci)
	defer os.Unsetenv("CI")
	result := applyCIMultiplier(timeout)
	if result != time.Duration(float64(timeout)*CITimeoutMultiplier) {
		t.Errorf("applyCIMultiplier(%v) = %v, want %v", timeout, result, time.Duration(float64(timeout)*CITimeoutMultiplier))
	}
}
