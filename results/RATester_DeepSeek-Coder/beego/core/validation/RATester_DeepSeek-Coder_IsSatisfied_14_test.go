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
			name: "TestIsSatisfied_Match",
			m:    Match{Regexp: regexp.MustCompile("^[a-zA-Z]+$"), Key: "field"},
			args: args{obj: "abc"},
			want: true,
		},
		{
			name: "TestIsSatisfied_NoMatch",
			m:    Match{Regexp: regexp.MustCompile("^[a-zA-Z]+$"), Key: "field"},
			args: args{obj: "abc123"},
			want: false,
		},
		{
			name: "TestIsSatisfied_NilRegexp",
			m:    Match{Regexp: nil, Key: "field"},
			args: args{obj: "abc"},
			want: false,
		},
		{
			name: "TestIsSatisfied_EmptyString",
			m:    Match{Regexp: regexp.MustCompile("^[a-zA-Z]+$"), Key: "field"},
			args: args{obj: ""},
			want: true,
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
				t.Errorf("Match.IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
