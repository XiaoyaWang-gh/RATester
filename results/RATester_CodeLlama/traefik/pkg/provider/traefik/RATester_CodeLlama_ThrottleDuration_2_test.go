package traefik

import (
	"fmt"
	"testing"
	"time"
)

func TestThrottleDuration_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	i := Provider{}
	want := time.Duration(0)
	got := i.ThrottleDuration()
	if got != want {
		t.Errorf("ThrottleDuration() = %v, want %v", got, want)
	}
}
