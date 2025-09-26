package glob

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterGlobParts_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "no glob",
			args: []string{"foo", "bar", "baz"},
			want: []string{"foo", "bar", "baz"},
		},
		{
			name: "with glob",
			args: []string{"foo", "*", "bar"},
			want: []string{"foo", "bar"},
		},
		{
			name: "empty",
			args: []string{},
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

			if got := FilterGlobParts(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FilterGlobParts() = %v, want %v", got, tt.want)
			}
		})
	}
}
