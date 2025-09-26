package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInclude_3(t *testing.T) {
	type args struct {
		cList []ControllerInterface
	}
	tests := []struct {
		name string
		n    *Namespace
		args args
		want *Namespace
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

			if got := tt.n.Include(tt.args.cList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Include() = %v, want %v", got, tt.want)
			}
		})
	}
}
