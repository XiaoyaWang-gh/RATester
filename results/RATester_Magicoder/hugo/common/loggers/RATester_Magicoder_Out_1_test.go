package loggers

import (
	"fmt"
	"io"
	"reflect"
	"strings"
	"testing"

	"github.com/bep/logg"
)

func TestOut_1(t *testing.T) {
	type fields struct {
		logCounters *logLevelCounter
		errors      *strings.Builder
		reset       func()
		out         io.Writer
		level       logg.Level
		logger      logg.Logger
		tracel      logg.LevelLogger
		debugl      logg.LevelLogger
		infol       logg.LevelLogger
		warnl       logg.LevelLogger
		errorl      logg.LevelLogger
	}
	tests := []struct {
		name   string
		fields fields
		want   io.Writer
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

			l := &logAdapter{
				logCounters: tt.fields.logCounters,
				errors:      tt.fields.errors,
				reset:       tt.fields.reset,
				out:         tt.fields.out,
				level:       tt.fields.level,
				logger:      tt.fields.logger,
				tracel:      tt.fields.tracel,
				debugl:      tt.fields.debugl,
				infol:       tt.fields.infol,
				warnl:       tt.fields.warnl,
				errorl:      tt.fields.errorl,
			}
			if got := l.Out(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logAdapter.Out() = %v, want %v", got, tt.want)
			}
		})
	}
}
