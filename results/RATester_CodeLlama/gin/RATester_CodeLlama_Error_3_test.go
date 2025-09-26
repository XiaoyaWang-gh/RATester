package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	msg := Error{Err: errors.New("test")}
	if msg.Error() != "test" {
		t.Errorf("Error() = %v, want %v", msg.Error(), "test")
	}
}
