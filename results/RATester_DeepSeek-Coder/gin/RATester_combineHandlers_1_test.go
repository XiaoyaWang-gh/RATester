package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCombineHandlers_1(t *testing.T) {
	type args struct {
		group    *RouterGroup
		handlers HandlersChain
	}
	tests := []struct {
		name string
		args args
		want HandlersChain
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

			if got := tt.args.group.combineHandlers(tt.args.handlers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combineHandlers() = %v, want %v", got, tt.want)
			}
		})
	}
}
