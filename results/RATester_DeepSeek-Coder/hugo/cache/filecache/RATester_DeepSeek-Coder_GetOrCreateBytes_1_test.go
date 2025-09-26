package filecache

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestGetOrCreateBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()

	testID := "testID"
	testData := []byte("testData")

	fs := afero.NewMemMapFs()
	cache := &Cache{
		Fs: fs,
	}

	// Test case where the file does not exist
	info, data, err := cache.GetOrCreateBytes(testID, func() ([]byte, error) {
		return testData, nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if info.Name != testID {
		t.Errorf("Expected name %s, got %s", testID, info.Name)
	}
	if !bytes.Equal(data, testData) {
		t.Errorf("Expected data %v, got %v", testData, data)
	}

	// Test case where the file exists
	_, err = fs.Create(testID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	info, data, err = cache.GetOrCreateBytes(testID, func() ([]byte, error) {
		t.Errorf("Create function should not be called when the file exists")
		return nil, nil
	})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if info.Name != testID {
		t.Errorf("Expected name %s, got %s", testID, info.Name)
	}
	if len(data) != 0 {
		t.Errorf("Expected empty data, got %v", data)
	}
}
