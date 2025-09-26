package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestfindCaseInsensitivePathRec_1(t *testing.T) {
	type args struct {
		path             string
		ciPath           []byte
		rb               [4]byte
		fixTrailingSlash bool
	}
	tests := []struct {
		name string
		n    *node
		args args
		want []byte
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

			if got := tt.n.findCaseInsensitivePathRec(tt.args.path, tt.args.ciPath, tt.args.rb, tt.args.fixTrailingSlash); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findCaseInsensitivePathRec() = %v, want %v", got, tt.want)
			}
		})
	}
}
