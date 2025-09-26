package echo

import (
	"fmt"
	"testing"
)

func TestAcquireContext_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	e := New()
	e.pool.New = func() interface{} {
		return new(Context)
	}

	ctx := e.AcquireContext()
	if ctx == nil {
		t.Errorf("Expected non-nil context, got nil")
	}

	e.ReleaseContext(ctx)
}
