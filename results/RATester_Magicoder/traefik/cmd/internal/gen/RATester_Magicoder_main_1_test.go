package main

import (
	"fmt"
	"go/build"
	"path"
	"path/filepath"
	"testing"
)

func TestMain_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dest := filepath.Join(path.Join(build.Default.GOPATH, "src"), destModuleName, destPkg)

	if dest != "github.com/traefik/genconf/dynamic" {
		t.Errorf("Expected dest to be 'github.com/traefik/genconf/dynamic', but got %s", dest)
	}

	err := run(dest)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
