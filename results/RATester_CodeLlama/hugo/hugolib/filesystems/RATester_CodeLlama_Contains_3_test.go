package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestContains_3(t *testing.T) {
	for _, test := range []struct {
		name     string
		filename string
		want     bool
	}{
		{"empty", "", false},
		{"no mounts", "foo", false},
		{"mounts", "foo/bar", true},
	} {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			d := &SourceFilesystem{
				Name: test.name,
				Fs:   afero.NewMemMapFs(),
			}
			if test.name == "mounts" {
				d.Fs.Mkdir("foo", 0755)
				d.Fs.Mkdir("foo/bar", 0755)
			}
			if got := d.Contains(test.filename); got != test.want {
				t.Errorf("Contains(%q) = %v, want %v", test.filename, got, test.want)
			}
		})
	}
}
