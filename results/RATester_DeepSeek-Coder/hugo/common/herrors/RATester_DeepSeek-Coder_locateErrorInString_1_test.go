package herrors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLocateErrorInString_1(t *testing.T) {
	type args struct {
		src     string
		matcher LineMatcherFn
	}
	tests := []struct {
		name string
		args args
		want *ErrorContext
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

			if got := locateErrorInString(tt.args.src, tt.args.matcher); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("locateErrorInString() = %v, want %v", got, tt.want)
			}
		})
	}
}
