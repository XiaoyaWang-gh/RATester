package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
)

func TestDebug_1(t *testing.T) {
	logger := zerolog.New(os.Stdout)
	l := OxyWrapper{logger: logger}

	tests := []struct {
		name string
		msg  string
		args []interface{}
	}{
		{
			name: "TestDebug",
			msg:  "This is a debug message",
			args: []interface{}{"arg1", "arg2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			l.Debug(tt.msg, tt.args...)
		})
	}
}
