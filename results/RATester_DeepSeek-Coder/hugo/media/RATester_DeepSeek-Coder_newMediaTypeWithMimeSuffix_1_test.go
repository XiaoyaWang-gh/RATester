package media

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewMediaTypeWithMimeSuffix_1(t *testing.T) {
	type args struct {
		main       string
		sub        string
		mimeSuffix string
		suffixes   []string
	}
	tests := []struct {
		name string
		args args
		want Type
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

			if got := newMediaTypeWithMimeSuffix(tt.args.main, tt.args.sub, tt.args.mimeSuffix, tt.args.suffixes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newMediaTypeWithMimeSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}
