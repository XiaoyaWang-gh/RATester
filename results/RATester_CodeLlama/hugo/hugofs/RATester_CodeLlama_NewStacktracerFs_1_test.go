package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestNewStacktracerFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	fs := afero.NewMemMapFs()
	pattern := ".*"

	// Act
	stacktracerFs := NewStacktracerFs(fs, pattern)

	// Assert
	assert.NotNil(t, stacktracerFs)
}
