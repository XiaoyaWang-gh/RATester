package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewResourceTransformationKey_1(t *testing.T) {
	type args struct {
		name     string
		elements []any
	}
	tests := []struct {
		name string
		args args
		want ResourceTransformationKey
	}{
		{
			name: "test",
			args: args{
				name:     "test",
				elements: []any{},
			},
			want: ResourceTransformationKey{
				Name:     "test",
				elements: []any{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := NewResourceTransformationKey(tt.args.name, tt.args.elements...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewResourceTransformationKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
