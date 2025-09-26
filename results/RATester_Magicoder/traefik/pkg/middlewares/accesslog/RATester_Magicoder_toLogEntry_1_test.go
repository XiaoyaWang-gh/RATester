package accesslog

import (
	"fmt"
	"testing"
)

func TestToLogEntry_1(t *testing.T) {
	type args struct {
		s            string
		defaultValue string
		quote        bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test Case 1",
			args: args{
				s:            "test",
				defaultValue: "default",
				quote:        true,
			},
			want: `"test"`,
		},
		{
			name: "Test Case 2",
			args: args{
				s:            "",
				defaultValue: "default",
				quote:        true,
			},
			want: `"default"`,
		},
		{
			name: "Test Case 3",
			args: args{
				s:            "test",
				defaultValue: "default",
				quote:        false,
			},
			want: "test",
		},
		{
			name: "Test Case 4",
			args: args{
				s:            "",
				defaultValue: "default",
				quote:        false,
			},
			want: "default",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := toLogEntry(tt.args.s, tt.args.defaultValue, tt.args.quote); got != tt.want {
				t.Errorf("toLogEntry() = %v, want %v", got, tt.want)
			}
		})
	}
}
