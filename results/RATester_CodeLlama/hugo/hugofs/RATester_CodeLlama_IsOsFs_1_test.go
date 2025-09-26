package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestIsOsFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewOsFs()
	if !IsOsFs(fs) {
		t.Errorf("IsOsFs() = %v, want %v", false, true)
	}
}
