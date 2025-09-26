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
				dst: Params{"key1": "value1", "key2": Params{"nestedKey1": "nestedValue1"}},
				src: Params{"key1": "newValue1", "key3": "value3"},
			},
			want: Params{"key1": "newValue1", "key2": Params{"nestedKey1": "nestedValue1"}, "key3": "value3"},
		},
		{
			name: "Test case 2",
			args: args{
				dst: Params{"key1": "value1", "key2": Params{"nestedKey1": "nestedValue1"}},
				src: Params{"key2": Params{"nestedKey1": "newNestedValue1", "nestedKey2": "nestedValue2"}},
			},
			want: Params{"key1": "value1", "key2": Params{"nestedKey1": "newNestedValue1", "nestedKey2": "nestedValue2"}},
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

			SetParams(tt.args.dst, tt.args.src)
			if !reflect.DeepEqual(tt.args.dst, tt.want) {
				t.Errorf("SetParams() = %v, want %v", tt.args.dst, tt.want)
			}
		})
	}
}
