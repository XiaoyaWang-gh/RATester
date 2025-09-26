package logs

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	kitlog "github.com/go-kit/log"
	"github.com/rs/zerolog"
)

func TestNewGoKitWrapper_1(t *testing.T) {
	type args struct {
		logger zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want kitlog.LoggerFunc
	}{
		{
			name: "test case 1",
			args: args{
				logger: zerolog.New(os.Stdout),
			},
			want: func(args ...interface{}) error {
				return nil
			},
		},
		{
			name: "test case 2",
			args: args{
				logger: zerolog.New(os.Stdout).With().Logger(),
			},
			want: func(args ...interface{}) error {
				return nil
			},
		},
		{
			name: "test case 3",
			args: args{
				logger: zerolog.New(os.Stdout).With().Logger().Level(zerolog.DebugLevel),
			},
			want: func(args ...interface{}) error {
				return nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewGoKitWrapper(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGoKitWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
