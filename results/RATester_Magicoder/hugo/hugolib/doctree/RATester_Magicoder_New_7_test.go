package doctree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNew_7(t *testing.T) {
	type args struct {
		cfg Config[any]
	}
	tests := []struct {
		name string
		args args
		want *NodeShiftTree[any]
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

			if got := New(tt.args.cfg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
