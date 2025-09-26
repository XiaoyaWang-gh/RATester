package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParamIfNotSet_1(t *testing.T) {
	type args struct {
		key   string
		value any
		d     *FrontMatterDescriptor
	}
	tests := []struct {
		name string
		args args
		want *FrontMatterDescriptor
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

			setParamIfNotSet(tt.args.key, tt.args.value, tt.args.d)
			if !reflect.DeepEqual(tt.args.d, tt.want) {
				t.Errorf("setParamIfNotSet() = %v, want %v", tt.args.d, tt.want)
			}
		})
	}
}
