package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMapBool_1(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name string
		args args
		want map[string]bool
	}{
		{
			name: "Test case 1",
			args: args{
				in: map[string]any{
					"key1": true,
					"key2": false,
				},
			},
			want: map[string]bool{
				"key1": true,
				"key2": false,
			},
		},
		{
			name: "Test case 2",
			args: args{
				in: map[string]any{
					"key3": "true",
					"key4": "false",
				},
			},
			want: map[string]bool{
				"key3": true,
				"key4": false,
			},
		},
		{
			name: "Test case 3",
			args: args{
				in: map[string]any{
					"key5": 1,
					"key6": 0,
				},
			},
			want: map[string]bool{
				"key5": true,
				"key6": false,
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

			if got := ToStringMapBool(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMapBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
