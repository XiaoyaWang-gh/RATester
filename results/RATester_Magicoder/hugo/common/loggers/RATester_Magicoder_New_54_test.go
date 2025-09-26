package loggers

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_54(t *testing.T) {
	type args struct {
		opts Options
	}
	tests := []struct {
		name string
		args args
		want Logger
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

			if got := New(tt.args.opts); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
