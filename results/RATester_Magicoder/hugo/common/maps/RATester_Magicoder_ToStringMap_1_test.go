package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStringMap_1(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "Test case 1",
			args: args{
				in: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
			},
			want: map[string]any{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "Test case 2",
			args: args{
				in: "string",
			},
			want: map[string]any{},
		},
		{
			name: "Test case 3",
			args: args{
				in: 123,
			},
			want: map[string]any{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := ToStringMap(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToStringMap() = %v, want %v", got, tt.want)
			}
		})
	}
}
