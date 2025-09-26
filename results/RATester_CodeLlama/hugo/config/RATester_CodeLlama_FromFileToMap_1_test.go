package config

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestFromFileToMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	fs := afero.NewMemMapFs()
	filename := "test.json"
	content := `{"key1": "value1", "key2": "value2"}`
	afero.WriteFile(fs, filename, []byte(content), 0644)
	// Act
	result, err := FromFileToMap(fs, filename)
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"key1": "value1", "key2": "value2"}, result)
}
