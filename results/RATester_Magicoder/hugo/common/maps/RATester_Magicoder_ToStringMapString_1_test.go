package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMapString_1(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "Test case 1",
			args: args{
				in: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Test case 2",
			args: args{
				in: "not a map",
			},
			want: map[string]string{},
		},
		{
			name: "Test case 3",
			args: args{
				in: nil,
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToStringMapString(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}
