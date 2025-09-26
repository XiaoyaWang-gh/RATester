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
			rule: "HostSNI:*",
			want: -1,
		},
		{
			name: "Test case 2",
			rule: "HostSNI:example.com",
			want: 17,
		},
		{
			name: "Test case 3",
			rule: "HostSNI:",
			want: 13,
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
