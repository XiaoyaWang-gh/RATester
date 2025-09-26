package filecache

import (
	"fmt"
	"testing"
)

func TestCacheDirModules_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	configs := Configs{
		CacheKeyModules: {
			DirCompiled: "/tmp/modules",
		},
	}

	expected := "/tmp/modules"
	actual := configs.CacheDirModules()

	if actual != expected {
		t.Errorf("Expected %s, got %s", expected, actual)
	}
}
