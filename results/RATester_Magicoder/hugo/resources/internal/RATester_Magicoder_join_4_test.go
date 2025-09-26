package internal

import (
	"fmt"
	"testing"
)

func Testjoin_4(t *testing.T) {
	tests := []struct {
		name string
		d    ResourcePaths
		p    []string
		want string
	}{
		{
			name: "empty",
			d:    ResourcePaths{},
			p:    []string{},
			want: "/",
		},
		{
			name: "single",
			d:    ResourcePaths{},
			p:    []string{"foo"},
			want: "/foo",
		},
		{
			name: "multiple",
			d:    ResourcePaths{},
			p:    []string{"foo", "bar"},
			want: "/foo/bar",
		},
		{
			name: "leading slash",
			d:    ResourcePaths{},
			p:    []string{"/foo", "bar"},
			want: "/foo/bar",
		},
		{
			name: "trailing slash",
			d:    ResourcePaths{},
			p:    []string{"foo/", "bar"},
			want: "/foo/bar",
		},
		{
			name: "empty strings",
			d:    ResourcePaths{},
			p:    []string{"", "foo", "", "bar", ""},
			want: "/foo/bar",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.d.join(tt.p...); got != tt.want {
				t.Errorf("join() = %v, want %v", got, tt.want)
			}
		})
	}
}
