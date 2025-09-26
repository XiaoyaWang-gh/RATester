package npm

import (
	"errors"
	"fmt"
	"testing"
)

func TestErr_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &packageBuilder{
		err: errors.New("test error"),
	}

	if err := b.Err(); err.Error() != "test error" {
		t.Errorf("Expected error to be 'test error', but got '%v'", err)
	}
}
