package maps

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetParams_2(t *testing.T) {
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
				dst: Params{"a": "1", "b": Params{"c": "2", "d": "3"}},
				src: Params{"a": "4", "b": Params{"c": "5", "e": "6"}},
			},
			want: Params{"a": "4", "b": Params{"c": "5", "d": "3", "e": "6"}},
		},
		{
			name: "Test case 2",
			args: args{
				dst: Params{"a": "1", "b": "2"},
				src: Params{"a": "4", "b": "5"},
			},
			want: Params{"a": "4", "b": "5"},
		},
		{
			name: "Test case 3",
			args: args{
				dst: Params{"a": "1", "b": Params{"c": "2", "d": "3"}},
				src: Params{"a": "4", "b": "5"},
			},
			want: Params{"a": "4", "b": "5"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			SetParams(tt.args.dst, tt.args.src)
			if !reflect.DeepEqual(tt.args.dst, tt.want) {
				t.Errorf("SetParams() = %v, want %v", tt.args.dst, tt.want)
			}
		})
	}
}
