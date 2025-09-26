package logs

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rs/zerolog"
)

func TestNewWasmLogger_1(t *testing.T) {
	type args struct {
		logger *zerolog.Logger
	}
	tests := []struct {
		name string
		args args
		want *WasmLogger
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

			if got := NewWasmLogger(tt.args.logger); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewWasmLogger() = %v, want %v", got, tt.want)
			}
		})
	}
}
