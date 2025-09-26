package context

import (
	"fmt"
	"reflect"
	"testing"
)

func TestbindPoint_1(t *testing.T) {
	type args struct {
		key string
		typ reflect.Type
	}
	tests := []struct {
		name  string
		input *BeegoInput
		args  args
		want  reflect.Value
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

			if got := tt.input.bindPoint(tt.args.key, tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bindPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}
