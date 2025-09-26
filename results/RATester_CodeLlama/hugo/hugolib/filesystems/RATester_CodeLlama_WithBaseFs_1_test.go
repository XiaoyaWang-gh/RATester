package filesystems

import (
	"fmt"
	"testing"
)

func TestWithBaseFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var b *BaseFs
	var bb *BaseFs
	var err error
	bb = &BaseFs{}
	err = WithBaseFs(b)(bb)
	if err != nil {
		t.Errorf("WithBaseFs(b)(bb) = %v, want nil", err)
	}
}
