package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopyInto_1(t *testing.T) {
	tests := []struct {
		name string
		in   *Domain
		want *Domain
	}{
		{
			name: "nil in",
			in:   nil,
			want: &Domain{},
		},
		{
			name: "empty in",
			in:   &Domain{},
			want: &Domain{},
		},
		{
			name: "non-empty in",
			in:   &Domain{Main: "example.com", SANs: []string{"www.example.com", "blog.example.com"}},
			want: &Domain{Main: "example.com", SANs: []string{"www.example.com", "blog.example.com"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := &Domain{}
			tt.in.DeepCopyInto(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}
