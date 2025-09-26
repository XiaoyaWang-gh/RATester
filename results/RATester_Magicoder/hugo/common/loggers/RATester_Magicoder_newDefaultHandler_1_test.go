package loggers

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestnewDefaultHandler_1(t *testing.T) {
	type args struct {
		outWriter io.Writer
		errWriter io.Writer
	}
	tests := []struct {
		name string
		args args
		want logg.Handler
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

			if got := newDefaultHandler(tt.args.outWriter, tt.args.errWriter); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newDefaultHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
