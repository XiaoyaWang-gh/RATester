package main

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
	"text/template"
)

func TestMain_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	os.Args = []string{"test", "linux"}
	main()

	if len(os.Args) < 2 {
		t.Error("GOOS should be provided as a CLI argument")
	}

	goos := strings.TrimSpace(os.Args[1])
	if goos == "" {
		t.Error("GOOS should be provided as a CLI argument")
	}

	tmpl := template.Must(
		template.New(".goreleaser.yml.tmpl").
			Delims("[[", "]]").
			ParseFiles("./.goreleaser.yml.tmpl"),
	)

	outputPath := path.Join(os.TempDir(), fmt.Sprintf(".goreleaser_%s.yml", goos))

	output, err := os.Create(outputPath)
	if err != nil {
		t.Error(err)
	}

	err = tmpl.Execute(output, map[string]string{"GOOS": goos})
	if err != nil {
		t.Error(err)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		t.Error("File does not exist")
	}
}
