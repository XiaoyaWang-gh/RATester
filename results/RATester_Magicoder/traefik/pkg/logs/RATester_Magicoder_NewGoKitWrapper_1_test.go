package logs

import (
	"fmt"
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
		// TODO: Add test cases.
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
