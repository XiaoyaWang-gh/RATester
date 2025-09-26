package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestNewFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	source := afero.NewMemMapFs()
	destination := afero.NewMemMapFs()
	workingDir := "."
	publishDir := "."

	// when
	fs := newFs(source, destination, workingDir, publishDir)

	// then
	assert.NotNil(t, fs)
	assert.Equal(t, source, fs.Source)
	assert.Equal(t, destination, fs.PublishDir)
	assert.Equal(t, destination, fs.PublishDirServer)
	assert.Equal(t, destination, fs.PublishDirStatic)
	assert.Equal(t, afero.NewOsFs(), fs.Os)
	assert.Equal(t, afero.NewMemMapFs(), fs.WorkingDirReadOnly)
	assert.Equal(t, afero.NewMemMapFs(), fs.WorkingDirWritable)
}
