package aggregator

import (
	"fmt"
	"testing"
)

func TestNewRingChannel_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ch := newRingChannel()
	if ch == nil {
		t.Error("newRingChannel() failed")
	}
}
