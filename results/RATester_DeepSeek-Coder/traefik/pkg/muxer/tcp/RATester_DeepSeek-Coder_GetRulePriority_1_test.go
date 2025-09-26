package tcp

import (
	"fmt"
	"testing"
)

func TestGetRulePriority_1(t *testing.T) {
	tests := []struct {
		name string
		rule string
		want int
	}{
		{
			name: "Test case 1",
			rule: "*",
			want: -1,
		},
		{
			name: "Test case 2",
			rule: "example.com",
			want: 12,
		},
		{
			name: "Test case 3",
			rule: "*:80",
			want: 5,
		},
		{
			name: "Test case 4",
			rule: "example.com:443",
			want: 16,
		},
		{
			name: "Test case 5",
			rule: "*:*",
			want: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetRulePriority(tt.rule); got != tt.want {
				t.Errorf("GetRulePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
