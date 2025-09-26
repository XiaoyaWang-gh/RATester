package warpc

import (
	"fmt"
	"testing"
	"time"
)

func TestputTimer_1(t *testing.T) {
	type args struct {
		t *time.Timer
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				t: time.NewTimer(time.Second),
			},
		},
		{
			name: "Test Case 2",
			args: args{
				t: time.NewTimer(time.Second),
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			putTimer(tt.args.t)
		})
	}
}
