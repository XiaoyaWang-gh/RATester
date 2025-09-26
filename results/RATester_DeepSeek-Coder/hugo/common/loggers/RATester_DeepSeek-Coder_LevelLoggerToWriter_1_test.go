package loggers

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestLevelLoggerToWriter_1(t *testing.T) {
	type args struct {
		l logg.LevelLogger
	}
	tests := []struct {
		name string
		args args
		want io.Writer
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

			if got := LevelLoggerToWriter(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelLoggerToWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
