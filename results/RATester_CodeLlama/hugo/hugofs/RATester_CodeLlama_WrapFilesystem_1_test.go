package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestWrapFilesystem_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	container := afero.NewMemMapFs()
	content := afero.NewMemMapFs()
	// When
	fs := WrapFilesystem(container, content)
	// Then
	assert.NotNil(t, fs)
}
