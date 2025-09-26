package deploy

import (
	"fmt"
	"reflect"
	"regexp"
	"testing"
)

func TestApplyOrdering_1(t *testing.T) {
	type args struct {
		ordering []*regexp.Regexp
		uploads  []*fileToUpload
	}
	tests := []struct {
		name string
		args args
		want [][]*fileToUpload
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

			if got := applyOrdering(tt.args.ordering, tt.args.uploads); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("applyOrdering() = %v, want %v", got, tt.want)
			}
		})
	}
}
