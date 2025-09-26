package herrors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewLineNumberErrHandlerFromRegexp_1(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want lineNumberExtractor
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

			got := newLineNumberErrHandlerFromRegexp(tt.args.expression)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newLineNumberErrHandlerFromRegexp() = %v, want %v", got, tt.want)
			}
		})
	}
}
