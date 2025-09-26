package filecache

import (
	"fmt"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestNewCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	fs := afero.NewMemMapFs()
	maxAge := time.Duration(10)
	pruneAllRootDir := "pruneAllRootDir"

	// when
	cache := NewCache(fs, maxAge, pruneAllRootDir)

	// then
	assert.NotNil(t, cache)
	assert.Equal(t, fs, cache.Fs)
	assert.NotNil(t, cache.nlocker)
	assert.Equal(t, maxAge, cache.maxAge)
	assert.Equal(t, pruneAllRootDir, cache.pruneAllRootDir)
}
