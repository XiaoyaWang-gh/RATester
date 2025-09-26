package loggers

import (
	"fmt"
	"testing"
	"time"
)

func TestPrintTimerIfDelayed_1(t *testing.T) {
	type args struct {
		start time.Time
		name  string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				start: time.Now(),
				name:  "Test",
			},
		},
		{
			name: "Test case 2",
			args: args{
				start: time.Now().Add(time.Second * 2),
				name:  "Test",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l := &logAdapter{}
			l.PrintTimerIfDelayed(tt.args.start, tt.args.name)
		})
	}
}
