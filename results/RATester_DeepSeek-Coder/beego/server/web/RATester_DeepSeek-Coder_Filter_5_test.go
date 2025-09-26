package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilter_5(t *testing.T) {
	type args struct {
		action string
		filter []FilterFunc
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

			if got := tt.n.Filter(tt.args.action, tt.args.filter...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Namespace.Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}
