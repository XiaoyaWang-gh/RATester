package htime

import (
	"fmt"
	"testing"
	"time"
)

func TestNow_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	want := time.Now()
	got := Now()
	if !got.Equal(want) {
		t.Errorf("Now() = %v, want %v", got, want)
	}
}
