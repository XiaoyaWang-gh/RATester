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

	if err := b.Err(); err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
