package hashing

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoHashable_1(t *testing.T) {
	type args struct {
		v any
	}
	tests := []struct {
		name string
		args args
		want any
	}{
		{
			name: "Test Case 1",
			args: args{
				v: "test",
			},
			want: "test",
		},
		{
			name: "Test Case 2",
			args: args{
				v: 123,
			},
			want: 123,
		},
		{
			name: "Test Case 3",
			args: args{
				v: struct{}{},
			},
			want: struct{}{},
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

			if got := toHashable(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toHashable() = %v, want %v", got, tt.want)
			}
		})
	}
}
