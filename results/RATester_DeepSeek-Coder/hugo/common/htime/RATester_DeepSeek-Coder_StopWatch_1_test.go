package htime

import (
	"fmt"
	"testing"
	"time"
)

func TestStopWatch_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "TestStopWatch"
	stop := StopWatch(name)
	time.Sleep(1 * time.Second)
	stop()
}
