package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TesthasHeaderValue_1(t *testing.T) {
	type args struct {
		m     http.Header
		key   string
		value string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test Case 1",
			args: args{
				m: http.Header{
					"Content-Type": []string{"application/json"},
				},
				key:   "Content-Type",
				value: "application/json",
			},
			want: true,
		},
		{
			name: "Test Case 2",
			args: args{
				m: http.Header{
					"Content-Type": []string{"application/json"},
				},
				key:   "Content-Type",
				value: "text/plain",
			},
			want: false,
		},
		{
			name: "Test Case 3",
			args: args{
				m: http.Header{
					"Content-Type": []string{"application/json"},
				},
				key:   "Content-Length",
				value: "application/json",
			},
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

			if got := hasHeaderValue(tt.args.m, tt.args.key, tt.args.value); got != tt.want {
				t.Errorf("hasHeaderValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
