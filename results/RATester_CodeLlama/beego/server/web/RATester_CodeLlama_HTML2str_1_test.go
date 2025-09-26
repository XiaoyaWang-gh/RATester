package web

import (
	"fmt"
	"testing"
)

func TestHTML2str_1(t *testing.T) {
	var tests = []struct {
		name string
		html string
		want string
	}{
		{
			name: "test1",
			html: "<html><head><title>test</title></head><body><h1>test</h1></body></html>",
			want: "test\ntest",
		},
		{
			name: "test2",
			html: "<html><head><title>test</title></head><body><h1>test</h1></body></html>",
			want: "test\ntest",
		},
		{
			name: "test3",
			html: "<html><head><title>test</title></head><body><h1>test</h1></body></html>",
			want: "test\ntest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HTML2str(tt.html); got != tt.want {
				t.Errorf("HTML2str() = %v, want %v", got, tt.want)
			}
		})
	}
}
