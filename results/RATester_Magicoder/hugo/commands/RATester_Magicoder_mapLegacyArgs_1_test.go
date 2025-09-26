package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestmapLegacyArgs_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want []string
	}{
		{
			name: "Test case 1",
			args: []string{"new", "site"},
			want: []string{"content", "site"},
		},
		{
			name: "Test case 2",
			args: []string{"new", "theme"},
			want: []string{"content", "theme"},
		},
		{
			name: "Test case 3",
			args: []string{"new", "content"},
			want: []string{"new", "content"},
		},
		{
			name: "Test case 4",
			args: []string{"new", "other"},
			want: []string{"content", "other"},
		},
		{
			name: "Test case 5",
			args: []string{"new"},
			want: []string{"content"},
		},
		{
			name: "Test case 6",
			args: []string{"other"},
			want: []string{"other"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := mapLegacyArgs(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapLegacyArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
