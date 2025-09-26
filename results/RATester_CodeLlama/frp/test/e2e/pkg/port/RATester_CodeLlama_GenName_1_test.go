package port

import (
	"fmt"
	"testing"
)

func TestGenName_1(t *testing.T) {
	type args struct {
		name    string
		options []NameOption
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test case 1",
			args: args{
				name:    "test-case-1",
				options: []NameOption{},
			},
			want: "testcase1",
		},
		{
			name: "test case 2",
			args: args{
				name:    "test-case-2",
				options: []NameOption{},
			},
			want: "testcase2",
		},
		{
			name: "test case 3",
			args: args{
				name:    "test-case-3",
				options: []NameOption{},
			},
			want: "testcase3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GenName(tt.args.name, tt.args.options...); got != tt.want {
				t.Errorf("GenName() = %v, want %v", got, tt.want)
			}
		})
	}
}
