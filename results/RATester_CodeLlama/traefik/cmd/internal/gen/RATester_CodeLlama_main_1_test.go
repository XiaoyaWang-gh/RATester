package main

import (
	"fmt"
	"go/build"
	"log"
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

	destModuleName := "github.com/golang/example"
	destPkg := "hello"

	dest := filepath.Join(path.Join(build.Default.GOPATH, "src"), destModuleName, destPkg)

	log.Println("Output:", dest)

	err := run(dest)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		t.Errorf("run() failed: %v", err)
	}
}
