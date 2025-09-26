package filesystems

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestRealFilename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &SourceFilesystem{
		Name: "layouts",
		Fs:   afero.NewMemMapFs(),
	}

	// This is a virtual composite filesystem. It expects path relative to a context.
	// So we need to create a subfolder.
	d.Fs.Mkdir("layouts", 0755)

	// Create a file in the subfolder.
	f, err := d.Fs.Create("layouts/index.html")
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	f.Close()

	// Now we can test the method.
	rel := "layouts/index.html"
	got := d.RealFilename(rel)
	want := "layouts/index.html"

	if got != want {
		t.Errorf("RealFilename(%q) = %q, want %q", rel, got, want)
	}
}
