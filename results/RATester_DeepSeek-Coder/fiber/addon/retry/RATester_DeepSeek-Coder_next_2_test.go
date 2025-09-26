package retry

import (
	"fmt"
	"testing"
	"time"
)

func TestNext_2(t *testing.T) {
	tests := []struct {
		name string
		e    *ExponentialBackoff
		want time.Duration
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.next(); got != tt.want {
				t.Errorf("ExponentialBackoff.next() = %v, want %v", got, tt.want)
			}
		})
	}
}
