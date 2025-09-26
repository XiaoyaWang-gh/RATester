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

	// 4.0.2
	// To upgrade to a new version, change to the commit hash of the version you want to upgrade to
	// then run mage generate from the root.
	const liveReloadCommit = "d803a41804d2d71e0814c4e9e3233e78991024d9"
	liveReloadSourceURL := fmt.Sprintf("https://raw.githubusercontent.com/livereload/livereload-js/%s/dist/livereload.js", liveReloadCommit)

	func() {
		resp, err := http.Get(liveReloadSourceURL)
		if err != nil {
			t.Fatal(err)
		}
		defer resp.Body.Close()

		b, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}

		err = os.WriteFile("../livereload.js", b, 0o644)
		if err != nil {
			t.Fatal(err)
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

		if len(result.Errors) > 0 {
			t.Fatal(result.Errors)
		}
	}()
}
