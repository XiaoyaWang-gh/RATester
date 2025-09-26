package hugofs

import (
	"fmt"
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
	}
	if fs.UnwrapFilesystem() != fs.Fs {
		t.Errorf("UnwrapFilesystem() = %v, want %v", fs.UnwrapFilesystem(), fs.Fs)
	}
}
