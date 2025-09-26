package metrics

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd_2(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		d    *diff
		args args
		want *diff
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

			if got := tt.d.add(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("diff.add() = %v, want %v", got, tt.want)
			}
		})
	}
}
