package filecache

import (
	"fmt"
	"testing"
)

func TestCleanID_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with relative path",
			args: args{name: "./test/../test"},
			want: "test",
		},
		{
			name: "Test with absolute path",
			args: args{name: "/test/../test"},
			want: "test",
		},
		{
			name: "Test with no path",
			args: args{name: "test"},
			want: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := cleanID(tt.args.name); got != tt.want {
				t.Errorf("cleanID() = %v, want %v", got, tt.want)
			}
		})
	}
}
