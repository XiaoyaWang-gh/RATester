package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestWalkFilesystems_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var fs afero.Fs
	var fn WalkFn

	// When
	var actual bool
	actual = WalkFilesystems(fs, fn)

	// Then
	if actual {
		t.Errorf("WalkFilesystems() = %v, want %v", actual, false)
	}
}
