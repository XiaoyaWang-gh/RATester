package hugofs

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

func TestDevOnly_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &stacktracerFs{
		Fs: afero.NewMemMapFs(),
		re: regexp.MustCompile(".*"),
	}

	fs.DevOnly()

	// Add assertions here
}
