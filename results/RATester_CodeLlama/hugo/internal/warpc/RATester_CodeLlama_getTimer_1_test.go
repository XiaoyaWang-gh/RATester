package warpc

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTimer_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := time.Duration(1)
	timer := getTimer(d)
	if timer == nil {
		t.Errorf("getTimer() = nil, want not nil")
	}
}
