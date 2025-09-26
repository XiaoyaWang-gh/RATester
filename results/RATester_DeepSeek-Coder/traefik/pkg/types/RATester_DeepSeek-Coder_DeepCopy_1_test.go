package types

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeepCopy_1(t *testing.T) {
	tests := []struct {
		name string
		in   *Domain
		want *Domain
	}{
		{
			name: "nil input",
			in:   nil,
			want: nil,
		},
		{
			name: "non-nil input",
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

			got := tt.in.DeepCopy()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DeepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}
