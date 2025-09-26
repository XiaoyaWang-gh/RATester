package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestGenKVDynConfDoc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	outputFile := "test.toml"
	defer os.Remove(outputFile)

	genKVDynConfDoc(outputFile)

	_, err := os.Stat(outputFile)
	if err != nil {
		t.Errorf("Expected file %s to exist, but it does not", outputFile)
	}

	content, err := ioutil.ReadFile(outputFile)
	if err != nil {
		t.Errorf("Error reading file %s: %v", outputFile, err)
	}

	if !strings.Contains(string(content), "| `traefik` | `") {
		t.Errorf("Expected file %s to contain key 'traefik', but it does not", outputFile)
	}
}
