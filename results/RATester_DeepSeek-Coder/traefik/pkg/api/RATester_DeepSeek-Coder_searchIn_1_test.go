package api

import (
	"fmt"
	"testing"
)

func TestSearchIn_1(t *testing.T) {
	type args struct {
		values []string
	}
	tests := []struct {
		name string
		c    *searchCriterion
		args args
		want bool
	}{
		{
			name: "should return true when search is empty",
			c:    &searchCriterion{Search: ""},
			args: args{values: []string{"test1", "test2", "test3"}},
			want: true,
		},
		{
			name: "should return true when search is found in values",
			c:    &searchCriterion{Search: "test1"},
			args: args{values: []string{"test1", "test2", "test3"}},
			want: true,
		},
		{
			name: "should return false when search is not found in values",
			c:    &searchCriterion{Search: "notfound"},
			args: args{values: []string{"test1", "test2", "test3"}},
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

			if got := tt.c.searchIn(tt.args.values...); got != tt.want {
				t.Errorf("searchCriterion.searchIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
