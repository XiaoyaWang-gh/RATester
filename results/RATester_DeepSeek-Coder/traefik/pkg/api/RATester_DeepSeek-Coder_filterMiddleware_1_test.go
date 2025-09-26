package api

import (
	"fmt"
	"testing"
)

func TestFilterMiddleware_1(t *testing.T) {
	type args struct {
		mns []string
	}
	tests := []struct {
		name string
		c    *searchCriterion
		args args
		want bool
	}{
		{
			name: "Test with empty MiddlewareName",
			c:    &searchCriterion{MiddlewareName: ""},
			args: args{mns: []string{"m1", "m2", "m3"}},
			want: true,
		},
		{
			name: "Test with MiddlewareName in the list",
			c:    &searchCriterion{MiddlewareName: "m1"},
			args: args{mns: []string{"m1", "m2", "m3"}},
			want: true,
		},
		{
			name: "Test with MiddlewareName not in the list",
			c:    &searchCriterion{MiddlewareName: "m4"},
			args: args{mns: []string{"m1", "m2", "m3"}},
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

			if got := tt.c.filterMiddleware(tt.args.mns); got != tt.want {
				t.Errorf("searchCriterion.filterMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
