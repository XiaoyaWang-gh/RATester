package loggers

import (
	"fmt"
	"io"
	"strings"
	"testing"

	"github.com/bep/logg"
)

func TestIdfInfoStatement_1(t *testing.T) {
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
	type args struct {
		what   string
		id     string
		format string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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
			if got := l.idfInfoStatement(tt.args.what, tt.args.id, tt.args.format); got != tt.want {
				t.Errorf("logAdapter.idfInfoStatement() = %v, want %v", got, tt.want)
			}
		})
	}
}
