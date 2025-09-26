package hugofs

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

func TestonCreate_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &stacktracerFs{
		Fs: afero.NewMemMapFs(),
		re: regexp.MustCompile(".*"),
	}

	filename := "test.txt"
	fs.onCreate(filename)

	// Check if the file exists
	_, err := fs.Fs.Stat(filename)
	if err != nil {
		t.Errorf("Expected file %s to exist, but it does not", filename)
	}
}
