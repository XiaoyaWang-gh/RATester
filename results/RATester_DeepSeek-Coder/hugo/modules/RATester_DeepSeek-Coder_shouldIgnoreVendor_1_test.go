package modules

import (
	"fmt"
	"testing"

	"github.com/gobwas/glob"
)

func TestShouldIgnoreVendor_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		c    ClientConfig
		args args
		want bool
	}{
		{
			name: "Test shouldIgnoreVendor with vendor directory",
			c:    ClientConfig{IgnoreVendor: glob.MustCompile("*_vendor")},
			args: args{path: "test_vendor"},
			want: true,
		},
		{
			name: "Test shouldIgnoreVendor without vendor directory",
			c:    ClientConfig{IgnoreVendor: glob.MustCompile("*_vendor")},
			args: args{path: "test"},
			want: false,
		},
		{
			name: "Test shouldIgnoreVendor with nil IgnoreVendor",
			c:    ClientConfig{IgnoreVendor: nil},
			args: args{path: "test_vendor"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.shouldIgnoreVendor(tt.args.path); got != tt.want {
				t.Errorf("ClientConfig.shouldIgnoreVendor() = %v, want %v", got, tt.want)
			}
		})
	}
}
