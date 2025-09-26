package loggers

import (
	"fmt"
	"testing"
	"time"

	"github.com/bep/logg"
)

func TestTimeTrackf_1(t *testing.T) {
	type args struct {
		l      logg.LevelLogger
		start  time.Time
		fields logg.Fields
		format string
		a      []any
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

			TimeTrackf(tt.args.l, tt.args.start, tt.args.fields, tt.args.format, tt.args.a...)
		})
	}
}
