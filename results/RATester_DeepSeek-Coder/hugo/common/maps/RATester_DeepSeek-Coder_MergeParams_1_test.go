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
				dst: Params{"a": "1", "b": "2", "c": "3"},
				src: Params{"b": "2.1", "c": "3.1", "d": "4"},
			},
			want: Params{"a": "1", "b": "2.1", "c": "3.1", "d": "4"},
		},
		{
			name: "Test case 2",
			args: args{
				dst: Params{"a": "1", "b": "2", "c": "3"},
				src: Params{"a": "1.1", "b": "2.1", "c": "3.1", "d": "4"},
			},
			want: Params{"a": "1.1", "b": "2.1", "c": "3.1", "d": "4"},
		},
		{
			name: "Test case 3",
			args: args{
				dst: Params{"a": "1", "b": "2", "c": "3"},
				src: Params{"a": "1", "b": "2", "c": "3"},
			},
			want: Params{"a": "1", "b": "2", "c": "3"},
		},
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
