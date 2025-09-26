package echo

import (
	"fmt"
	"testing"
	"time"
)

func TestUnixTimeNano_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	binder := &ValueBinder{
		ValueFunc: func(sourceParam string) string {
			if sourceParam == "test" {
				return "1609459200000000000"
			}
			return ""
		},
	}

	var dest time.Time
	err := binder.UnixTimeNano("test", &dest)

	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	expected := time.Unix(1609459200, 0)
	if !dest.Equal(expected) {
		t.Errorf("Expected %v, but got: %v", expected, dest)
	}
}
