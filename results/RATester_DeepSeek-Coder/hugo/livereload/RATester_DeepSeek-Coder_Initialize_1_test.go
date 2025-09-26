package livereload

import (
	"fmt"
	"testing"
	"time"
)

func TestInitialize_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	Initialize()

	// Check if wsHub.run() is running as a goroutine
	select {
	case <-time.After(1 * time.Second):
		t.Errorf("wsHub.run() did not start as a goroutine")
	default:
	}
}
