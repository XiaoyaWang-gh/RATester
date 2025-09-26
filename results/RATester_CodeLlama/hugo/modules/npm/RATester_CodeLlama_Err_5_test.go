package npm

import (
	"fmt"
	"testing"
)

func TestErr_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &packageBuilder{}
	if b.Err() != nil {
		t.Errorf("expected nil, got %v", b.Err())
	}
}
