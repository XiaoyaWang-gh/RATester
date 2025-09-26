package filecache

import (
	"fmt"
	"testing"

	"github.com/BurntSushi/locker"
	"github.com/spf13/afero"
)

func TestWriteCloser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	fs := afero.NewMemMapFs()
	c := &Cache{
		Fs: fs,
		nlocker: &lockTracker{
			seen:   make(map[string]struct{}),
			Locker: &locker.Locker{},
		},
	}

	testID := "testID"
	expectedName := cleanID(testID)

	info, wc, err := c.WriteCloser(testID)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if info.Name != expectedName {
		t.Errorf("Expected name %q, got %q", expectedName, info.Name)
	}

	if wc == nil {
		t.Error("Expected a non-nil io.WriteCloser, got nil")
	}

	_, err = fs.Stat(expectedName)
	if err != nil {
		t.Errorf("Expected file %q to exist, got error: %v", expectedName, err)
	}

	wc.Close()

	_, err = fs.Stat(expectedName)
	if err == nil {
		t.Errorf("Expected file %q to not exist, but it does", expectedName)
	}
}
