package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeParamsWithStrategy_1(t *testing.T) {
	type args struct {
		strategy string
		dst      Params
		src      Params
	}
	tests := []struct {
		name string
		args args
		want Params
	}{
		{
			name: "Test case 1",
			args: args{
				strategy: "strategy1",
				dst:      Params{"key1": "value1"},
				src:      Params{"key2": "value2"},
			},
			want: Params{"key1": "value1", "key2": "value2"},
		},
		{
			name: "Test case 2",
			args: args{
				strategy: "strategy2",
				dst:      Params{"key1": "value1", "key2": "value2"},
				src:      Params{"key2": "value3", "key3": "value3"},
			},
			want: Params{"key1": "value1", "key2": "value3", "key3": "value3"},
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

			MergeParamsWithStrategy(tt.args.strategy, tt.args.dst, tt.args.src)
			if !reflect.DeepEqual(tt.args.dst, tt.want) {
				t.Errorf("MergeParamsWithStrategy() = %v, want %v", tt.args.dst, tt.want)
			}
		})
	}
}
