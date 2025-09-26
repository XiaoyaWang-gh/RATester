package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestMakePathRelative_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	var (
		d = &SourceFilesystem{
			Name: "layouts",
			Fs:   afero.NewMemMapFs(),
		}
		filename    = "layouts/index.html"
		checkExists = true
	)

	if _, ok := d.MakePathRelative(filename, checkExists); !ok {
		t.Errorf("MakePathRelative(%q, %t) = %q, %t; want %q, %t", filename, checkExists, "", false, filename, true)
	}
}
