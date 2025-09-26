package logs

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
)

func TestRun_2(t *testing.T) {
	type args struct {
		e     *zerolog.Event
		level zerolog.Level
		msg   string
	}
	tests := []struct {
		name string
		n    NoLevelHook
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

			tt.n.Run(tt.args.e, tt.args.level, tt.args.msg)
		})
	}
}
