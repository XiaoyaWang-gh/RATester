package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestErrorf_1(t *testing.T) {
	logger := ElasticLogger{
		logger: zerolog.New(os.Stdout),
	}

	tests := []struct {
		name   string
		format string
		args   []interface{}
	}{
		{
			name:   "TestErrorf",
			format: "This is an error message",
			args:   []interface{}{"arg1", "arg2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			logger.Errorf(tt.format, tt.args...)
		})
	}
}
