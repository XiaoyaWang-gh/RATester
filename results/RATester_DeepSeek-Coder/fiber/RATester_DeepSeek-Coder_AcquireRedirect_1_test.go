package fiber

import (
	"fmt"
	"testing"
)

func TestAcquireRedirect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	redirect := AcquireRedirect()

	if redirect == nil {
		t.Errorf("Expected a non-nil Redirect, got nil")
	}

	redirect.release()
}
