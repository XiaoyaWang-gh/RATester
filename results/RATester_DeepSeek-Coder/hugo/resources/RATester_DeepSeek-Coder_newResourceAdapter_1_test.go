package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewResourceAdapter_1(t *testing.T) {
	t.Parallel()

	type args struct {
		spec        *Spec
		lazyPublish bool
		target      transformableResource
	}

	tests := []struct {
		name string
		args args
		want *resourceAdapter
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

			got := newResourceAdapter(tt.args.spec, tt.args.lazyPublish, tt.args.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newResourceAdapter() = %v, want %v", got, tt.want)
			}
		})
	}
}
