package gin

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestLoggerWithWriter_1(t *testing.T) {
	type args struct {
		out       io.Writer
		notlogged []string
	}
	tests := []struct {
		name string
		args args
		want HandlerFunc
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

			if got := LoggerWithWriter(tt.args.out, tt.args.notlogged...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoggerWithWriter() = %v, want %v", got, tt.want)
			}
		})
	}
}
