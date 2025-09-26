package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeParams_1(t *testing.T) {
	type args struct {
		dst Params
		src Params
	}
	tests := []struct {
		name string
		args args
		want Params
	}{
		{
			name: "Test case 1",
			args: args{
				dst: Params{"key1": "value1", "key2": "value2"},
				src: Params{"key3": "value3", "key4": "value4"},
			},
			want: Params{"key1": "value1", "key2": "value2", "key3": "value3", "key4": "value4"},
		},
		{
			name: "Test case 2",
			args: args{
				dst: Params{"key1": "value1", "key2": "value2"},
				src: Params{"key1": "newValue1", "key3": "value3"},
			},
			want: Params{"key1": "newValue1", "key2": "value2", "key3": "value3"},
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

			MergeParams(tt.args.dst, tt.args.src)
			if !reflect.DeepEqual(tt.args.dst, tt.want) {
				t.Errorf("MergeParams() = %v, want %v", tt.args.dst, tt.want)
			}
		})
	}
}
