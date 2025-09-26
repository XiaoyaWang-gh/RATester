package hstring

import (
	"fmt"
	"testing"
)

func TestString_3(t *testing.T) {
	type fields struct {
		HTML HTML
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test case 1",
			fields: fields{
				HTML: "<p>Hello, World</p>",
			},
			want: "<p>Hello, World</p>",
		},
		{
			name: "Test case 2",
			fields: fields{
				HTML: "<p>Hello, <a href=\"https://example.com\">World</a></p>",
			},
			want: "<p>Hello, <a href=\"https://example.com\">World</a></p>",
		},
		{
			name: "Test case 3",
			fields: fields{
				HTML: "<p>Hello, <a href=\"https://example.com\">World</a></p>",
			},
			want: "<p>Hello, <a href=\"https://example.com\">World</a></p>",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := HTML(tt.fields.HTML)
			if got := s.String(); got != tt.want {
				t.Errorf("HTML.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
