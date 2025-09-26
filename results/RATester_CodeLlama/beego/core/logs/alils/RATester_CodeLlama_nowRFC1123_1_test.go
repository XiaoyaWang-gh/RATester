package alils

import (
	"fmt"
	"testing"
	"time"
)

func TestNowRFC1123_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	now := time.Now()
	got := nowRFC1123()
	want := now.In(gmtLoc).Format(time.RFC1123)
	if got != want {
		t.Errorf("nowRFC1123() = %v, want %v", got, want)
	}
}
