package skeletons

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestNewSiteCreateConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	fs := afero.NewMemMapFs()
	createpath := "/tmp/hugo"
	format := "yaml"

	// When
	err := newSiteCreateConfig(fs, createpath, format)

	// Then
	if err != nil {
		t.Errorf("newSiteCreateConfig() error = %v", err)
		return
	}

	// Check if the file exists
	if _, err := fs.Stat(filepath.Join(createpath, "hugo."+format)); os.IsNotExist(err) {
		t.Errorf("newSiteCreateConfig() file does not exist")
	}
}
