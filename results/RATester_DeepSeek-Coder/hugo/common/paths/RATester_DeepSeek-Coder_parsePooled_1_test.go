package paths

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParsePooled_1(t *testing.T) {
	type args struct {
		c string
		s string
	}
	tests := []struct {
		name string
		pp   *PathParser
		args args
		want *Path
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

			if got := tt.pp.parsePooled(tt.args.c, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PathParser.parsePooled() = %v, want %v", got, tt.want)
			}
		})
	}
}
