package glob

import (
	"fmt"
	"testing"
)

func TestAppend_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &FilenameFilter{}
	other := &FilenameFilter{}
	result := f.Append(other)
	if result == nil {
		t.Error("result is nil")
	}
	if len(result.nested) != 2 {
		t.Errorf("len(result.nested) = %d, want 2", len(result.nested))
	}
	if result.nested[0] != f {
		t.Errorf("result.nested[0] = %p, want %p", result.nested[0], f)
	}
	if result.nested[1] != other {
		t.Errorf("result.nested[1] = %p, want %p", result.nested[1], other)
	}
}
