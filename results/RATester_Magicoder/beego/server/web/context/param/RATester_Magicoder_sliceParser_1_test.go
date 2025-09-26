package param

import (
	"fmt"
	"reflect"
	"testing"
)

func TestsliceParser_1(t *testing.T) {
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

			if got := sliceParser(tt.args.elemParser); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sliceParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
