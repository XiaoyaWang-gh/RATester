package hugolib

import (
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/frankban/quicktest"
	"github.com/spf13/afero"
)

func TestwriteToFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Create a new instance of IntegrationTestBuilder
	builder := &IntegrationTestBuilder{
		C: &quicktest.C{},
	}

	// Create a new instance of afero.MemMapFs
	fs := afero.NewMemMapFs()

	// Define the filename and content for the file
	filename := "test.txt"
	content := "This is a test file."

	// Call the method under test
	builder.writeToFs(fs, filename, content)

	// Check if the file was created and has the correct content
	file, err := fs.Open(filename)
	if err != nil {
		t.Errorf("Failed to open file: %s", err)
	}
	defer file.Close()

	data, err := ioutil.ReadAll(file)
	if err != nil {
		t.Errorf("Failed to read file: %s", err)
	}

	if string(data) != content {
		t.Errorf("Expected content: %s, got: %s", content, string(data))
	}
}
