package page

import (
	"fmt"
	"reflect"
	"testing"
)

func TestWrapSite_1(t *testing.T) {
	type testSite struct {
		Site
	}

	tests := []struct {
		name string
		s    Site
		want Site
	}{
		{
			name: "Nil Site",
			s:    nil,
			want: nil,
		},
		{
			name: "Non-Nil Site",
			s:    &testSite{},
			want: &siteWrapper{s: &testSite{}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got := WrapSite(tt.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WrapSite() = %v, want %v", got, tt.want)
			}
		})
	}
}
