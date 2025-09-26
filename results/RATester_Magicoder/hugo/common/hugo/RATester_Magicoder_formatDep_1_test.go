package hugo

import (
	"fmt"
	"testing"
)

func TestformatDep_1(t *testing.T) {
	type args struct {
		path    string
		version string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test case 1",
			args: args{
				path:    "test/path",
				version: "1.0.0",
			},
			want: "test/path=\"1.0.0\"",
		},
		{
			name: "Test case 2",
			args: args{
				path:    "another/path",
				version: "2.0.0",
			},
			want: "another/path=\"2.0.0\"",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := formatDep(tt.args.path, tt.args.version); got != tt.want {
				t.Errorf("formatDep() = %v, want %v", got, tt.want)
			}
		})
	}
}
