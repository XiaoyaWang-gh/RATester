package task

import (
	"fmt"
	"testing"
	"time"
)

func TestDayMatches_1(t *testing.T) {
	type args struct {
		s *Schedule
		t time.Time
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := dayMatches(tt.args.s, tt.args.t); got != tt.want {
				t.Errorf("dayMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
