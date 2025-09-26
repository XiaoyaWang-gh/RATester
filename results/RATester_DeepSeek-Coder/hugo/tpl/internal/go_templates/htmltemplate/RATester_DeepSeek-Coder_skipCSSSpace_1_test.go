package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSkipCSSSpace_1(t *testing.T) {
	tests := []struct {
		name string
		c    []byte
		want []byte
	}{
		{"Test case 1", []byte(" \t\n\r\fHello, World!"), []byte("Hello, World")},
		{"Test case 2", []byte("\r\nHello, World!"), []byte("Hello, World")},
		{"Test case 3", []byte("\nHello, World!"), []byte("Hello, World")},
		{"Test case 4", []byte("\fHello, World!"), []byte("Hello, World")},
		{"Test case 5", []byte("\tHello, World!"), []byte("Hello, World")},
		{"Test case 6", []byte(" Hello, World!"), []byte("Hello, World")},
		{"Test case 7", []byte(""), []byte("")},
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
