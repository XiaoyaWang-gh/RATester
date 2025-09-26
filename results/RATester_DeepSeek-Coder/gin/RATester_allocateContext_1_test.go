package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAllocateContext_1(t *testing.T) {
	type args struct {
		maxParams uint16
	}
	tests := []struct {
		name string
		args args
		want *Context
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

			engine := &Engine{}
			if got := engine.allocateContext(tt.args.maxParams); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("allocateContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
