package service

import (
	"fmt"
	"testing"
)

func TestcontainsNTLMorNegotiate_1(t *testing.T) {
	tests := []struct {
		name string
		h    []string
		want bool
	}{
		{
			name: "Test case 1",
			h:    []string{"NTLM", "Negotiate", "Basic"},
			want: true,
		},
		{
			name: "Test case 2",
			h:    []string{"Basic", "Bearer"},
			want: false,
		},
		{
			name: "Test case 3",
			h:    []string{"NTLM", "Negotiate", "Bearer"},
			want: true,
		},
		{
			name: "Test case 4",
			h:    []string{"Basic", "NTLM", "Negotiate"},
			want: true,
		},
		{
			name: "Test case 5",
			h:    []string{"Bearer", "NTLM", "Negotiate"},
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

			if got := containsNTLMorNegotiate(tt.h); got != tt.want {
				t.Errorf("containsNTLMorNegotiate() = %v, want %v", got, tt.want)
			}
		})
	}
}
