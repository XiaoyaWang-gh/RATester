package api

import (
	"fmt"
	"testing"
)

func TestFilterService_1(t *testing.T) {
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
			name: "Test with empty service name",
			c:    &searchCriterion{ServiceName: ""},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test with service name contains @",
			c:    &searchCriterion{ServiceName: "test@"},
			args: args{name: "test@"},
			want: true,
		},
		{
			name: "Test with service name does not contain @",
			c:    &searchCriterion{ServiceName: "test"},
			args: args{name: "test"},
			want: true,
		},
		{
			name: "Test with service name contains @ and name does not match",
			c:    &searchCriterion{ServiceName: "test@"},
			args: args{name: "notMatch"},
			want: false,
		},
		{
			name: "Test with service name does not contain @ and name matches",
			c:    &searchCriterion{ServiceName: "test"},
			args: args{name: "test"},
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
				t.Errorf("searchCriterion.filterService() = %v, want %v", got, tt.want)
			}
		})
	}
}
