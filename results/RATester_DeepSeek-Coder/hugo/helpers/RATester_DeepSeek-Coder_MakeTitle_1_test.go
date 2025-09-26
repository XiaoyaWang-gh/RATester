package helpers

import (
	"fmt"
	"testing"
)

func TestMakeTitle_1(t *testing.T) {
	type args struct {
		inpath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test with hyphen",
			args: args{inpath: "test-title"},
			want: "test title",
		},
		{
			name: "Test with multiple hyphens",
			args: args{inpath: "test----title"},
			want: "test  title",
		},
		{
			name: "Test with spaces",
			args: args{inpath: "test  title"},
			want: "test title",
		},
		{
			name: "Test with spaces and hyphens",
			args: args{inpath: "test - - title"},
			want: "test  title",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := MakeTitle(tt.args.inpath); got != tt.want {
				t.Errorf("MakeTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
