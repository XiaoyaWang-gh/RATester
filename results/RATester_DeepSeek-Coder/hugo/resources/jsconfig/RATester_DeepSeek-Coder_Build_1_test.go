package jsconfig

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"sync"
	"testing"
)

func TestBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Builder{
		sourceRootsMu: sync.RWMutex{},
		sourceRoots:   map[string]bool{"testdata/src": true},
	}

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	conf := b.Build(dir)
	if conf == nil {
		t.Fatal("expected a config, got nil")
	}

	expectedRoots := []string{"testdata/src"}
	sort.Strings(expectedRoots)

	roots := conf.CompilerOptions.Paths["*"]
	sort.Strings(roots)

	if !reflect.DeepEqual(roots, expectedRoots) {
		t.Errorf("expected roots %v, got %v", expectedRoots, roots)
	}
}
