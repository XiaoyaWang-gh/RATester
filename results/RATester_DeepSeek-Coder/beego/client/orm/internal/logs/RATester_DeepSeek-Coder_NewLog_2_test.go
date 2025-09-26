package logs

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewLog_2(t *testing.T) {
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

			if got := NewLog(tt.args.out); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
