package logs

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewElasticLogger_1(t *testing.T) {
	type args struct {
		logger zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want *ElasticLogger
	}{
		{
			name: "test case 1",
			args: args{
				logger: zerolog.New(os.Stdout),
			},
			want: &ElasticLogger{
				logger: zerolog.New(os.Stdout),
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

			if got := NewElasticLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewElasticLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
