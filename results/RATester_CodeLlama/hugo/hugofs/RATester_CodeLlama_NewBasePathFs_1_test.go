package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestNewBasePathFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	source := afero.NewMemMapFs()
	fs := NewBasePathFs(source, "test")
	assert.Equal(t, "test", fs.Name())
}
