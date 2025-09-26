package hugofs

import (
	"fmt"
	"os"
	"regexp"
	"testing"

	"github.com/spf13/afero"
)

func TestOpenFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &stacktracerFs{
		Fs: afero.NewMemMapFs(),
		re: regexp.MustCompile(".*"),
	}

	testFile := "test.txt"
	_, err := fs.OpenFile(testFile, os.O_CREATE, 0666)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = fs.Fs.Stat(testFile)
	if err != nil {
		t.Errorf("Expected file to exist, got error: %v", err)
	}

	fs.re = regexp.MustCompile("non-matching-pattern")
	_, err = fs.OpenFile(testFile, os.O_CREATE, 0666)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	_, err = fs.Fs.Stat(testFile)
	if err != nil {
		t.Errorf("Expected file to still exist, got error: %v", err)
	}
}
