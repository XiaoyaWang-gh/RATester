package logs

import (
	"fmt"
	"os"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/require"
)

func TestFatalln_1(t *testing.T) {
	logger := zerolog.New(os.Stdout)
	l := LogrusStdWrapper{logger: logger}

	tests := []struct {
		name string
		args []interface{}
	}{
		{
			name: "TestFatalln with one argument",
			args: []interface{}{"test"},
		},
		{
			name: "TestFatalln with multiple arguments",
			args: []interface{}{"test1", "test2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			require := require.New(t)
			require.Panics(func() { l.Fatalln(tt.args...) })
		})
	}
}
