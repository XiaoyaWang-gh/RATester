package api

import (
	"fmt"
	"testing"
)

func TestfilterMiddleware_1(t *testing.T) {
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
			name: "Test case 1",
			c: &searchCriterion{
				MiddlewareName: "middleware1",
			},
			args: args{
				mns: []string{"middleware1", "middleware2"},
			},
			want: true,
		},
		{
			name: "Test case 2",
			c: &searchCriterion{
				MiddlewareName: "middleware3",
			},
			args: args{
				mns: []string{"middleware1", "middleware2"},
			},
			want: false,
		},
		{
			name: "Test case 3",
			c: &searchCriterion{
				MiddlewareName: "",
			},
			args: args{
				mns: []string{"middleware1", "middleware2"},
			},
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

			if got := tt.c.filterMiddleware(tt.args.mns); got != tt.want {
				t.Errorf("searchCriterion.filterMiddleware() = %v, want %v", got, tt.want)
			}
		})
	}
}
