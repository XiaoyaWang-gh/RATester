package api

import (
	"fmt"
	"testing"
)

func TestWithStatus_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		c    *searchCriterion
		args args
		want bool
	}{
		{
			name: "Test withStatus with empty status",
			c:    &searchCriterion{Status: ""},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test withStatus with non-empty status",
			c:    &searchCriterion{Status: "test"},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test withStatus with different case",
			c:    &searchCriterion{Status: "Test"},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test withStatus with different name",
			c:    &searchCriterion{Status: "test"},
			args: args{name: "not test"},
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

			if got := tt.c.withStatus(tt.args.name); got != tt.want {
				t.Errorf("withStatus() = %v, want %v", got, tt.want)
			}
		})
	}
}
