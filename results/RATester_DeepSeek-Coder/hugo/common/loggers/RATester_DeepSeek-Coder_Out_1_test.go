package loggers

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestOut_1(t *testing.T) {
	tests := []struct {
		name string
		l    *logAdapter
		want io.Writer
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

			if got := tt.l.Out(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("logAdapter.Out() = %v, want %v", got, tt.want)
			}
		})
	}
}
