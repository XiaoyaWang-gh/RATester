package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestgetScheme_1(t *testing.T) {
	tests := []struct {
		name string
		uri  []byte
		want []byte
	}{
		{
			name: "Test case 1",
			uri:  []byte("http://example.com"),
			want: []byte("http"),
		},
		{
			name: "Test case 2",
			uri:  []byte("https://example.com"),
			want: []byte("https"),
		},
		{
			name: "Test case 3",
			uri:  []byte("ftp://example.com"),
			want: []byte("ftp"),
		},
		{
			name: "Test case 4",
			uri:  []byte("invalid"),
			want: nil,
		},
		{
			name: "Test case 5",
			uri:  []byte("http:/example.com"),
			want: nil,
		},
		{
			name: "Test case 6",
			uri:  []byte("http://"),
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getScheme(tt.uri); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getScheme() = %v, want %v", got, tt.want)
			}
		})
	}
}
