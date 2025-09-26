package logs

import (
	"fmt"
	"testing"

	"github.com/rs/zerolog"
)

func TestLogWithLevel_1(t *testing.T) {
	type args struct {
		ev  *zerolog.Event
		msg string
		kvs []interface{}
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

			logWithLevel(tt.args.ev, tt.args.msg, tt.args.kvs...)
		})
	}
}
