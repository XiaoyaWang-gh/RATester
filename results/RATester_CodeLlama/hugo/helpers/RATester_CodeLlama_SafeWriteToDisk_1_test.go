package helpers

import (
	"fmt"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

func TestSafeWriteToDisk_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	inpath := "test.txt"
	r := strings.NewReader("Hello World")
	fs := afero.NewMemMapFs()
	// Act
	err := SafeWriteToDisk(inpath, r, fs)
	// Assert
	if err != nil {
		t.Errorf("SafeWriteToDisk() error = %v", err)
		return
	}
	if _, err := fs.Stat(inpath); err != nil {
		t.Errorf("SafeWriteToDisk() file not found = %v", err)
		return
	}
}
