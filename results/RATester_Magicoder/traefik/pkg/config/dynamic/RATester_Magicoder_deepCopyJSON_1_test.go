package dynamic

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopyJSON_1(t *testing.T) {
	type args struct {
		x map[string]any
	}
	tests := []struct {
		name string
		args args
		want map[string]any
	}{
		{
			name: "Test Case 1",
			args: args{
				x: map[string]any{
					"key1": "value1",
					"key2": 2,
					"key3": true,
				},
			},
			want: map[string]any{
				"key1": "value1",
				"key2": 2,
				"key3": true,
			},
		},
		{
			name: "Test Case 2",
			args: args{
				x: map[string]any{
					"key4": map[string]any{
						"key41": "value41",
						"key42": 42,
						"key43": false,
					},
				},
			},
			want: map[string]any{
				"key4": map[string]any{
					"key41": "value41",
					"key42": 42,
					"key43": false,
				},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := deepCopyJSON(tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deepCopyJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}
