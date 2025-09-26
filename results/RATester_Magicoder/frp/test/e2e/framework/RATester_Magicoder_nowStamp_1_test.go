package framework

import (
	"fmt"
	"testing"
	"time"
)

func TestNowStamp_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := nowStamp()
	_, err := time.Parse(time.StampMilli, now)
	if err != nil {
		t.Errorf("nowStamp() = %s, should be a valid time stamp", now)
	}
}
