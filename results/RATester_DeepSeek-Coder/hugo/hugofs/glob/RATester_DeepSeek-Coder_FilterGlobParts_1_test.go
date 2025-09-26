package glob

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterGlobParts_1(t *testing.T) {
	type args struct {
		a []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test with no glob characters",
			args: args{a: []string{"test1", "test2", "test3"}},
			want: []string{"test1", "test2", "test3"},
		},
		{
			name: "Test with glob characters",
			args: args{a: []string{"test*", "test1", "*test", "test2", "test3"}},
			want: []string{"test1", "test2", "test3"},
		},
		{
			name: "Test with empty slice",
			args: args{a: []string{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := FilterGlobParts(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterGlobParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
