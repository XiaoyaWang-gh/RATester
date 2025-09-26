package skeletons

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestCreateSite_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	createpath := "path"
	sourceFs := afero.NewMemMapFs()
	force := false
	format := "yaml"

	// When
	err := CreateSite(createpath, sourceFs, force, format)

	// Then
	if err != nil {
		t.Errorf("CreateSite() error = %v", err)
		return
	}
}
