package hugofs

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

func TestUnwrapFilesystem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &stacktracerFs{
		Fs: afero.NewMemMapFs(),
		re: regexp.MustCompile(".*"),
	}

	expected := afero.NewMemMapFs()
	actual := fs.UnwrapFilesystem()

	if actual != expected {
		t.Errorf("Expected %v, but got %v", expected, actual)
	}
}
