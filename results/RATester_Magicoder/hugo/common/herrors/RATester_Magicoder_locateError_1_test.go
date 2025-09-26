package herrors

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestlocateError_1(t *testing.T) {
	type args struct {
		r       io.Reader
		le      FileError
		matches LineMatcherFn
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

			if got := locateError(tt.args.r, tt.args.le, tt.args.matches); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("locateError() = %v, want %v", got, tt.want)
			}
		})
	}
}
