package filecache

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/spf13/afero"
)

func TestNewCache_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	maxAge := time.Duration(10)
	pruneAllRootDir := "/tmp"

	cache := NewCache(fs, maxAge, pruneAllRootDir)

	if cache.Fs != fs {
		t.Errorf("Expected Fs to be %v, but got %v", fs, cache.Fs)
	}

	if cache.maxAge != maxAge {
		t.Errorf("Expected maxAge to be %v, but got %v", maxAge, cache.maxAge)
	}

	if cache.pruneAllRootDir != pruneAllRootDir {
		t.Errorf("Expected pruneAllRootDir to be %v, but got %v", pruneAllRootDir, cache.pruneAllRootDir)
	}

	if cache.nlocker == nil {
		t.Errorf("Expected nlocker to be not nil")
	}

	if cache.initOnce != (sync.Once{}) {
		t.Errorf("Expected initOnce to be %v, but got %v", sync.Once{}, cache.initOnce)
	}

	if cache.initErr != nil {
		t.Errorf("Expected initErr to be nil, but got %v", cache.initErr)
	}
}
