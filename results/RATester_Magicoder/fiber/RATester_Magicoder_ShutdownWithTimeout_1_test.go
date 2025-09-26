package fiber

import (
	"fmt"
	"testing"
	"time"
)

func TestShutdownWithTimeout_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	app := &App{}

	timeout := 10 * time.Second
	err := app.ShutdownWithTimeout(timeout)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
}
