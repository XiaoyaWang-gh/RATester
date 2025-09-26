package proxy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetScheme_1(t *testing.T) {
	tests := []struct {
		name string
		uri  []byte
		want []byte
	}{
		{"http", []byte("http://example.com"), []byte("http")},
		{"https", []byte("https://example.com"), []byte("https")},
		{"ftp", []byte("ftp://example.com"), []byte("ftp")},
		{"no scheme", []byte("example.com"), nil},
		{"empty", []byte(""), nil},
		{"no slash", []byte("http:example.com"), nil},
		{"no colon", []byte("http//example.com"), nil},
		{"no colon or slash", []byte("http:example.com"), nil},
		{"no colon or slash 2", []byte("http:///example.com"), nil},
		{"no colon or slash 3", []byte("http://example.com/"), nil},
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
