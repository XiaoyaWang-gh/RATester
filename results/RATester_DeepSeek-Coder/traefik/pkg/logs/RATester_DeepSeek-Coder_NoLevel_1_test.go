package logs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestNoLevel_1(t *testing.T) {
	type args struct {
		logger zerolog.Logger
		level  zerolog.Level
	}
	tests := []struct {
		name string
		args args
		want zerolog.Logger
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

			if got := NoLevel(tt.args.logger, tt.args.level); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NoLevel() = %v, want %v", got, tt.want)
			}
		})
	}
}
