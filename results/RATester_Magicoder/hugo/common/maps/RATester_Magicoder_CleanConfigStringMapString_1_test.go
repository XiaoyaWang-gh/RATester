package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCleanConfigStringMapString_1(t *testing.T) {
	tests := []struct {
		name string
		args map[string]string
		want map[string]string
	}{
		{
			name: "empty map",
			args: map[string]string{},
			want: map[string]string{},
		},
		{
			name: "map with MergeStrategyKey",
			args: map[string]string{
				MergeStrategyKey: "merge",
				"key1":           "value1",
				"key2":           "value2",
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		{
			name: "map without MergeStrategyKey",
			args: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			want: map[string]string{
				"key1": "value1",
				"key2": "value2",
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

			if got := CleanConfigStringMapString(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CleanConfigStringMapString() = %v, want %v", got, tt.want)
			}
		})
	}
}
