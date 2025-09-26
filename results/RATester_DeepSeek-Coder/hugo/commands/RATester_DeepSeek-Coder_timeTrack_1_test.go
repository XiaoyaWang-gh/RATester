package commands

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeTrack_1(t *testing.T) {
	type args struct {
		start time.Time
		name  string
	}
	tests := []struct {
		name string
		args args
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

			r := &rootCommand{}
			r.timeTrack(tt.args.start, tt.args.name)
		})
	}
}
