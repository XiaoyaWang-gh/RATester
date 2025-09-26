package alils

import (
	"fmt"
	"testing"
	"time"
)

func TestnowRFC1123_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := nowRFC1123()
	_, err := time.Parse(time.RFC1123, now)
	if err != nil {
		t.Errorf("nowRFC1123() = %v, want valid RFC1123 time", now)
	}
}
