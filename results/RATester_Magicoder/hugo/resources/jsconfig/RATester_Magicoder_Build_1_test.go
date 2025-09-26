package jsconfig

import (
	"fmt"
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
		sourceRoots:   map[string]bool{},
	}

	b.AddSourceRoot("/path/to/root1")
	b.AddSourceRoot("/path/to/root2")

	conf := b.Build("/path/to/dir")

	if conf == nil {
		t.Fatal("expected config, got nil")
	}

	if len(conf.CompilerOptions.Paths) != 1 {
		t.Fatalf("expected 1 path, got %d", len(conf.CompilerOptions.Paths))
	}

	paths, ok := conf.CompilerOptions.Paths["*"]
	if !ok {
		t.Fatal("expected path '*' to exist")
	}

	if len(paths) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(paths))
	}

	if paths[0] != "root1" || paths[1] != "root2" {
		t.Fatalf("expected paths to be ['root1', 'root2'], got %v", paths)
	}
}
