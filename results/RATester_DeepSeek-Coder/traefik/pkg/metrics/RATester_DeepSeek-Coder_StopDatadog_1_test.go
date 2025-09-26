package metrics

import (
	"fmt"
	"testing"
)

func TestStopDatadog_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	datadogLoopCancelFunc = func() {
		t.Log("datadogLoopCancelFunc was called")
	}

	StopDatadog()

	if datadogLoopCancelFunc != nil {
		t.Errorf("Expected datadogLoopCancelFunc to be nil, but it's not")
	}
}
