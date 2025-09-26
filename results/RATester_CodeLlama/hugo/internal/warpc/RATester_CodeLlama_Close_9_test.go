package warpc

import (
	"fmt"
	"testing"
)

func TestClose_9(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &Dispatchers{}
	if err := d.Close(); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
