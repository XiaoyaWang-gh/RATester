package jsconfig

import (
	"fmt"
	"testing"
)

func TestBuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	b := &Builder{}
	b.AddSourceRoot("/path/to/source")
	b.AddSourceRoot("/path/to/source2")
	conf := b.Build("/path/to/dir")
	if conf == nil {
		t.Fatal("expected a config")
	}
	if len(conf.CompilerOptions.Paths) != 1 {
		t.Fatalf("expected 1 path, got %d", len(conf.CompilerOptions.Paths))
	}
	if len(conf.CompilerOptions.Paths["*"]) != 2 {
		t.Fatalf("expected 2 paths, got %d", len(conf.CompilerOptions.Paths["*"]))
	}
	if conf.CompilerOptions.Paths["*"][0] != "/path/to/source/*" {
		t.Fatalf("expected %q, got %q", "/path/to/source/*", conf.CompilerOptions.Paths["*"][0])
	}
	if conf.CompilerOptions.Paths["*"][1] != "/path/to/source2/*" {
		t.Fatalf("expected %q, got %q", "/path/to/source2/*", conf.CompilerOptions.Paths["*"][1])
	}
}
