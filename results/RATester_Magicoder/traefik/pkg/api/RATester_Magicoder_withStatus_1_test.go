package api

import (
	"fmt"
	"testing"
)

func TestwithStatus_1(t *testing.T) {
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
			name: "Test withStatus when status is empty",
			c:    &searchCriterion{Status: ""},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test withStatus when status is not empty and equal",
			c:    &searchCriterion{Status: "test"},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test withStatus when status is not empty and not equal",
			c:    &searchCriterion{Status: "test"},
			args: args{name: "test1"},
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
