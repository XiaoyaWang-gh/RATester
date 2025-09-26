package loggers

import (
	"fmt"
	"io"
	"reflect"
	"testing"

	"github.com/bep/logg"
)

func TestNewNoColoursHandler_1(t *testing.T) {
	type args struct {
		outWriter     io.Writer
		errWriter     io.Writer
		noLevelPrefix bool
		predicate     func(*logg.Entry) bool
	}
	tests := []struct {
		name string
		args args
		want *noColoursHandler
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

			if got := newNoColoursHandler(tt.args.outWriter, tt.args.errWriter, tt.args.noLevelPrefix, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newNoColoursHandler() = %v, want %v", got, tt.want)
			}
		})
	}
}
