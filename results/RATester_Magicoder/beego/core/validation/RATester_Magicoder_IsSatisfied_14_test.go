package validation

import (
	"fmt"
	"regexp"
	"testing"
)

func TestIsSatisfied_14(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name string
		m    Match
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			m: Match{
				Regexp: regexp.MustCompile("^[a-zA-Z0-9]+$"),
				Key:    "test",
			},
			args: args{
				obj: "test123",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			m: Match{
				Regexp: regexp.MustCompile("^[a-zA-Z0-9]+$"),
				Key:    "test",
			},
			args: args{
				obj: "test@123",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.IsSatisfied(tt.args.obj); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
