package msg

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewDispatcher_1(t *testing.T) {
	type args struct {
		rw io.ReadWriter
	}
	tests := []struct {
		name string
		args args
		want *Dispatcher
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

			if got := NewDispatcher(tt.args.rw); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDispatcher() = %v, want %v", got, tt.want)
			}
		})
	}
}
