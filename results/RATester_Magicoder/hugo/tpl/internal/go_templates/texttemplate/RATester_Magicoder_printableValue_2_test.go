package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestprintableValue_2(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name  string
		args  args
		want  any
		want1 bool
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

			got, got1 := printableValue(tt.args.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("printableValue() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("printableValue() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
