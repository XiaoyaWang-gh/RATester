package gin

import (
	"fmt"
	"testing"
)

func TestIsType_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err Error
	err.Type = ErrorType(1)
	if !err.IsType(ErrorType(1)) {
		t.Errorf("IsType() = %v, want %v", err.IsType(ErrorType(1)), true)
	}
}
