package cache

import (
	"fmt"
	"testing"
)

func TestStartAndGC_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{}
	config := `{"CachePath":"/tmp/cache", "FileSuffix":".cache", "DirectoryLevel":"2", "EmbedExpiry":"3600"}`
	err := fc.StartAndGC(config)
	if err != nil {
		t.Errorf("StartAndGC failed, err: %v", err)
	}
	if fc.CachePath != "/tmp/cache" {
		t.Errorf("CachePath not set correctly, expected: /tmp/cache, got: %s", fc.CachePath)
	}
	if fc.FileSuffix != ".cache" {
		t.Errorf("FileSuffix not set correctly, expected: .cache, got: %s", fc.FileSuffix)
	}
	if fc.DirectoryLevel != 2 {
		t.Errorf("DirectoryLevel not set correctly, expected: 2, got: %d", fc.DirectoryLevel)
	}
	if fc.EmbedExpiry != 3600 {
		t.Errorf("EmbedExpiry not set correctly, expected: 3600, got: %d", fc.EmbedExpiry)
	}
}
