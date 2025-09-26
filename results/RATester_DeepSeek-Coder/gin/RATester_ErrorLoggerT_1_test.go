package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestErrorLoggerT_1(t *testing.T) {
	type args struct {
		typ ErrorType
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

			if got := ErrorLoggerT(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ErrorLoggerT() = %v, want %v", got, tt.want)
			}
		})
	}
}
