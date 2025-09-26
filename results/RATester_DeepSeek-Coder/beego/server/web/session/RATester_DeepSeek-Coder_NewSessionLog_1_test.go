package session

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewSessionLog_1(t *testing.T) {
	type args struct {
		out io.Writer
	}
	tests := []struct {
		name string
		args args
		want *Log
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

			if got := NewSessionLog(tt.args.out); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSessionLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
