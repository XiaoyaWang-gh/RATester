package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIsTextUnmarshaler_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want unmarshaler
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

			if got := isTextUnmarshaler(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("isTextUnmarshaler() = %v, want %v", got, tt.want)
			}
		})
	}
}
