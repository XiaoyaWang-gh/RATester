package httplib

import (
	"fmt"
	"testing"
)

func TestRetries_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &BeegoHTTPRequest{}
	times := 3
	b.Retries(times)
	if b.setting.Retries != times {
		t.Errorf("Expected Retries to be %d, but got %d", times, b.setting.Retries)
	}
}
