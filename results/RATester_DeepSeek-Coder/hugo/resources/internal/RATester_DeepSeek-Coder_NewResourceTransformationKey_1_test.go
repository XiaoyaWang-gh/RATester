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
			name: "Test case 1",
			args: args{
				name:     "testName",
				elements: []any{1, "test", 3.14},
			},
			want: ResourceTransformationKey{
				Name:     "testName",
				elements: []any{1, "test", 3.14},
			},
		},
		{
			name: "Test case 2",
			args: args{
				name:     "anotherTest",
				elements: []any{true, "another", 6.28},
			},
			want: ResourceTransformationKey{
				Name:     "anotherTest",
				elements: []any{true, "another", 6.28},
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
