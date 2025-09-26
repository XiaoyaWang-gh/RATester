package resource

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetParamToLower_1(t *testing.T) {
	type args struct {
		r   Resource
		key string
	}
	tests := []struct {
		name string
		args args
		want any
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

			if got := GetParamToLower(tt.args.r, tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetParamToLower() = %v, want %v", got, tt.want)
			}
		})
	}
}
