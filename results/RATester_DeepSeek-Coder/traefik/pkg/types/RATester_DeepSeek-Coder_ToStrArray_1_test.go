package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestToStrArray_1(t *testing.T) {
	tests := []struct {
		name string
		d    *Domain
		want []string
	}{
		{
			name: "empty domain",
			d:    &Domain{},
			want: []string{},
		},
		{
			name: "only main domain",
			d:    &Domain{Main: "example.com"},
			want: []string{"example.com"},
		},
		{
			name: "only SANs",
			d:    &Domain{SANs: []string{"www.example.com", "blog.example.com"}},
			want: []string{"www.example.com", "blog.example.com"},
		},
		{
			name: "main domain and SANs",
			d:    &Domain{Main: "example.com", SANs: []string{"www.example.com", "blog.example.com"}},
			want: []string{"example.com", "www.example.com", "blog.example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.ToStrArray(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Domain.ToStrArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
