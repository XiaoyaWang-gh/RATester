package htime

import (
	"fmt"
	"testing"
)

func TestStopWatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stop := StopWatch("TestStopWatch")
	defer stop()

	// Your test code here
}
