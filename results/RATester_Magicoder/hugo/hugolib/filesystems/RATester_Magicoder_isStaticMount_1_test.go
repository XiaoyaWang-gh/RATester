package filesystems

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/modules"
)

func TestisStaticMount_1(t *testing.T) {
	b := &sourceFilesystemsBuilder{}

	tests := []struct {
		name string
		mnt  modules.Mount
		want bool
	}{
		{
			name: "static mount",
			mnt: modules.Mount{
				Target: "/static/",
			},
			want: true,
		},
		{
			name: "non-static mount",
			mnt: modules.Mount{
				Target: "/assets/",
			},
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

			if got := b.isStaticMount(tt.mnt); got != tt.want {
				t.Errorf("isStaticMount() = %v, want %v", got, tt.want)
			}
		})
	}
}
