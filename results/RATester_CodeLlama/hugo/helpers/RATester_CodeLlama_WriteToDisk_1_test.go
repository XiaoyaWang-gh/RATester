package helpers

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/spf13/afero"
)

func TestWriteToDisk_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	inpath := "test.txt"
	r := strings.NewReader("Hello, World!")
	fs := afero.NewMemMapFs()
	// Act
	err := WriteToDisk(inpath, r, fs)
	// Assert
	if err != nil {
		t.Errorf("WriteToDisk() error = %v", err)
		return
	}
	if _, err := fs.Stat(inpath); os.IsNotExist(err) {
		t.Errorf("WriteToDisk() file not found: %s", inpath)
	}
}
