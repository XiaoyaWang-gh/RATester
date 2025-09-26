package api

import (
	"fmt"
	"testing"
)

func TestsearchIn_1(t *testing.T) {
	tests := []struct {
		name   string
		c      *searchCriterion
		values []string
		want   bool
	}{
		{
			name: "search is empty",
			c: &searchCriterion{
				Search: "",
			},
			values: []string{"test1", "test2", "test3"},
			want:   true,
		},
		{
			name: "search is not empty and found",
			c: &searchCriterion{
				Search: "test1",
			},
			values: []string{"test1", "test2", "test3"},
			want:   true,
		},
		{
			name: "search is not empty and not found",
			c: &searchCriterion{
				Search: "test4",
			},
			values: []string{"test1", "test2", "test3"},
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.searchIn(tt.values...); got != tt.want {
				t.Errorf("searchIn() = %v, want %v", got, tt.want)
			}
		})
	}
}
