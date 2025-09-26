package npm

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestPack_1(t *testing.T) {

	sourceFs := afero.NewMemMapFs()
	assetsWithDuplicatesPreservedFs := afero.NewMemMapFs()

	// TODO: add test cases.
	t.Run("test case 1", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		if err := Pack(sourceFs, assetsWithDuplicatesPreservedFs); err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
