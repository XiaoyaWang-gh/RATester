package commands

import (
	"fmt"
	"testing"
	"time"
)

func TestInitMemTicker_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &hugoBuilder{}

	// Call the method under test
	stop := h.initMemTicker()

	// Wait for the ticker to fire
	time.Sleep(6 * time.Second)

	// Call the stop function
	stop()

	// Add assertions here if needed
}
