package log

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWithLog_1(t *testing.T) {
	type args struct {
		f func(f interface{}, v ...interface{})
	}
	tests := []struct {
		name string
		args args
		want BuilderOption
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

			if got := WithLog(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WithLog() = %v, want %v", got, tt.want)
			}
		})
	}
}
