package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNSInclude_1(t *testing.T) {
	type args struct {
		cList []ControllerInterface
	}
	tests := []struct {
		name string
		args args
		want LinkNamespace
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

			if got := NSInclude(tt.args.cList...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NSInclude() = %v, want %v", got, tt.want)
			}
		})
	}
}
