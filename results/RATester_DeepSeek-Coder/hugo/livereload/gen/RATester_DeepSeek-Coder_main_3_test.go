package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/evanw/esbuild/pkg/api"
)

func TestMain_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	liveReloadCommit := "d803a41804d2d71e0814c4e9e3233e78991024d9"
	liveReloadSourceURL := fmt.Sprintf("https://raw.githubusercontent.com/livereload/livereload-js/%s/dist/livereload.js", liveReloadCommit)

	// Act
	resp, err := http.Get(liveReloadSourceURL)
	if err != nil {
		t.Errorf("Failed to get response from %s: %v", liveReloadSourceURL, err)
	}
	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("Failed to read response body: %v", err)
	}

	err = os.WriteFile("../livereload.js", b, 0o644)
	if err != nil {
		t.Errorf("Failed to write file: %v", err)
	}

	result := api.Build(api.BuildOptions{
		Stdin: &api.StdinOptions{
			Contents: string(b) + livereloadHugoPluginJS,
		},
		Outfile:           "../livereload.min.js",
		Bundle:            true,
		Target:            api.ES2015,
		Write:             true,
		MinifyWhitespace:  true,
		MinifyIdentifiers: true,
		MinifySyntax:      true,
	})

	// Assert
	if len(result.Errors) > 0 {
		t.Errorf("Build failed with errors: %v", result.Errors)
	}
}
