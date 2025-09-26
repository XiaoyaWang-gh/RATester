package commands

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMapLegacyArgs_1(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test with 'new' and 'site'",
			args: args{args: []string{"new", "site"}},
			want: []string{"new", "content", "site"},
		},
		{
			name: "Test with 'new' and 'theme'",
			args: args{args: []string{"new", "theme"}},
			want: []string{"new", "content", "theme"},
		},
		{
			name: "Test with 'new' and 'content'",
			args: args{args: []string{"new", "content"}},
			want: []string{"new", "content"},
		},
		{
			name: "Test with 'site'",
			args: args{args: []string{"site"}},
			want: []string{"site"},
		},
		{
			name: "Test with 'new' and 'other'",
			args: args{args: []string{"new", "other"}},
			want: []string{"new", "content", "other"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := mapLegacyArgs(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapLegacyArgs() = %v, want %v", got, tt.want)
			}
		})
	}
}
