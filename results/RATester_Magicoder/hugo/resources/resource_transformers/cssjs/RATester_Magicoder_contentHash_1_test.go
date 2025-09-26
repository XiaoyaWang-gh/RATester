package cssjs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestcontentHash_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	imp := &importResolver{
		fs: afero.NewMemMapFs(),
	}

	filename := "test.txt"
	content := "test content"
	err := afero.WriteFile(imp.fs, filename, []byte(content), 0644)
	if err != nil {
		t.Fatal(err)
	}

	b, hash := imp.contentHash(filename)

	if string(b) != content {
		t.Errorf("Expected content to be %s, but got %s", content, string(b))
	}

	expectedHash := "9f0df396d973748ac2202315f4588ecb2d1c45b2d1f7c4b8a9c1a07cd8113063"
	if hash != expectedHash {
		t.Errorf("Expected hash to be %s, but got %s", expectedHash, hash)
	}
}
