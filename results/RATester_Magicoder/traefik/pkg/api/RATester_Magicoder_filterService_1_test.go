package api

import (
	"fmt"
	"testing"
)

func TestfilterService_1(t *testing.T) {
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
			name: "Test case 1",
			c: &searchCriterion{
				ServiceName: "test@example.com",
			},
			args: args{
				name: "test",
			},
			want: true,
		},
		{
			name: "Test case 2",
			c: &searchCriterion{
				ServiceName: "test@example.com",
			},
			args: args{
				name: "test2",
			},
			want: false,
		},
		{
			name: "Test case 3",
			c: &searchCriterion{
				ServiceName: "test",
			},
			args: args{
				name: "test",
			},
			want: true,
		},
		{
			name: "Test case 4",
			c: &searchCriterion{
				ServiceName: "test@example.com",
			},
			args: args{
				name: "test@example.com",
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

			if got := tt.c.filterService(tt.args.name); got != tt.want {
				t.Errorf("filterService() = %v, want %v", got, tt.want)
			}
		})
	}
}
