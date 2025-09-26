package pagemeta

import (
	"fmt"
	"reflect"
	"testing"
)

func TesttoLowerSlice_1(t *testing.T) {
	type args struct {
		in any
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{
				in: []string{"HELLO", "WORLD"},
			},
			want: []string{"hello", "world"},
		},
		{
			name: "Test Case 2",
			args: args{
				in: []string{"Golang", "is", "Awesome"},
			},
			want: []string{"golang", "is", "awesome"},
		},
		{
			name: "Test Case 3",
			args: args{
				in: []string{"123", "456"},
			},
			want: []string{"123", "456"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toLowerSlice(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("toLowerSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
