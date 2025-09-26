package httplib

import (
	"fmt"
	"testing"
	"time"
)

func TestWithRetry_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	times := 1
	delay := time.Duration(1)
	WithRetry(times, delay)
}
