package npm

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestnewPackageBuilder_1(t *testing.T) {
	type args struct {
		source string
		first  io.Reader
	}
	tests := []struct {
		name string
		args args
		want *packageBuilder
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

			if got := newPackageBuilder(tt.args.source, tt.args.first); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newPackageBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
