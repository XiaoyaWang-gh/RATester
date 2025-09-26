package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestptrParser_1(t *testing.T) {
	type args struct {
		elemParser paramParser
	}
	tests := []struct {
		name string
		args args
		want paramParser
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

			if got := ptrParser(tt.args.elemParser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ptrParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
