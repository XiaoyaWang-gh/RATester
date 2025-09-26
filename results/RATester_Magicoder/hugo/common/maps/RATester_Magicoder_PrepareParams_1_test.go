package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestPrepareParams_1(t *testing.T) {
	tests := []struct {
		name string
		args Params
		want Params
	}{
		{
			name: "Test case 1",
			args: Params{
				"Key1": "Value1",
				"Key2": "Value2",
			},
			want: Params{
				"key1": "Value1",
				"key2": "Value2",
			},
		},
		{
			name: "Test case 2",
			args: Params{
				"Key1": "Value1",
				"Key2": Params{
					"Key3": "Value3",
					"Key4": "Value4",
				},
			},
			want: Params{
				"key1": "Value1",
				"key2": Params{
					"key3": "Value3",
					"key4": "Value4",
				},
			},
		},
		{
			name: "Test case 3",
			args: Params{
				"Key1": "Value1",
				"Key2": map[string]string{
					"Key3": "Value3",
					"Key4": "Value4",
				},
			},
			want: Params{
				"key1": "Value1",
				"key2": Params{
					"key3": "Value3",
					"key4": "Value4",
				},
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

			PrepareParams(tt.args)
			if !reflect.DeepEqual(tt.args, tt.want) {
				t.Errorf("PrepareParams() = %v, want %v", tt.args, tt.want)
			}
		})
	}
}
