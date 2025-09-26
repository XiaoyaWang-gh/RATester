package skeletons

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/helpers"
	"github.com/spf13/afero"
)

func TestCreateTheme_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	createpath := "test"
	sourceFs := afero.NewMemMapFs()
	// Act
	err := CreateTheme(createpath, sourceFs)
	// Assert
	if err != nil {
		t.Errorf("CreateTheme() error = %v", err)
		return
	}
	if exists, _ := helpers.Exists(createpath, sourceFs); !exists {
		t.Errorf("CreateTheme() error = %v", err)
		return
	}
}
