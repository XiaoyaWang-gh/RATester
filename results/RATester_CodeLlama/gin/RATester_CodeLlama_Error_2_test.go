package gin

import (
	"errors"
	"fmt"
	"testing"
)

func TestError_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &Context{}
	c.Error(errors.New("test error"))
	if len(c.Errors) != 1 {
		t.Errorf("Expected 1 error, got %d", len(c.Errors))
	}
}
