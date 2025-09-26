package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestskipCSSSpace_1(t *testing.T) {
	tests := []struct {
		name string
		c    []byte
		want []byte
	}{
		{
			name: "Test case 1",
			c:    []byte("\tHello"),
			want: []byte("Hello"),
		},
		{
			name: "Test case 2",
			c:    []byte("\nHello"),
			want: []byte("Hello"),
		},
		{
			name: "Test case 3",
			c:    []byte("\fHello"),
			want: []byte("Hello"),
		},
		{
			name: "Test case 4",
			c:    []byte(" Hello"),
			want: []byte("Hello"),
		},
		{
			name: "Test case 5",
			c:    []byte("\rHello"),
			want: []byte("Hello"),
		},
		{
			name: "Test case 6",
			c:    []byte("\r\nHello"),
			want: []byte("Hello"),
		},
		{
			name: "Test case 7",
			c:    []byte("Hello"),
			want: []byte("Hello"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := skipCSSSpace(tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("skipCSSSpace() = %v, want %v", got, tt.want)
			}
		})
	}
}
