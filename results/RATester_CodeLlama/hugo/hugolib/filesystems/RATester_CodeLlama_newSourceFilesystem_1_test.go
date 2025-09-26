package filesystems

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestNewSourceFilesystem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	fs := afero.NewMemMapFs()
	b := &sourceFilesystemsBuilder{}
	result := b.newSourceFilesystem(name, fs)
	assert.Equal(t, name, result.Name)
	assert.Equal(t, fs, result.Fs)
}
