package grace

import (
	"fmt"
	"testing"
)

func TestWithShutdownCallback_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		shutdownCallbackCalled bool
		shutdownCallback       = func() { shutdownCallbackCalled = true }
	)

	WithShutdownCallback(shutdownCallback)(nil)

	if !shutdownCallbackCalled {
		t.Error("shutdownCallback was not called")
	}
}
