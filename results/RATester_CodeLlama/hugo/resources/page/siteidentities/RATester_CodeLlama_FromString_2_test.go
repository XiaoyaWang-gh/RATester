package siteidentities

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/identity"
)

func TestFromString_2(t *testing.T) {
	tests := []struct {
		name  string
		want  identity.Identity
		want1 bool
	}{
		{
			name:  "Data",
			want:  Data,
			want1: true,
		},
		{
			name:  "Data",
			want:  identity.Anonymous,
			want1: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, got1 := FromString(tt.name)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("FromString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
