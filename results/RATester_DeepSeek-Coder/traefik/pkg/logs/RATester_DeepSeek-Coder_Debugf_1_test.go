package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestDebugf_1(t *testing.T) {
	logger := zerolog.New(os.Stdout)
	l := ElasticLogger{logger: logger}

	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{
			name:   "Test with string format and args",
			format: "This is a test %s",
			args:   []interface{}{"message"},
		},
		{
			name:   "Test with int format and args",
			format: "This is a test %d",
			args:   []interface{}{123},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l.Debugf(tt.format, tt.args...)
		})
	}
}
