package framework

import (
	"fmt"
	"os"
	"testing"
)

func TestGenerateConfigFile_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := NewFramework(Options{
		TotalParallelNode: 1,
		CurrentNodeIndex:  0,
		FromPortIndex:     10000,
		ToPortIndex:       20000,
	})
	defer f.AfterEach()

	f.BeforeEach()

	content := "test content"
	path := f.GenerateConfigFile(content)

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	if string(data) != content {
		t.Fatalf("Expected content: %s, got: %s", content, string(data))
	}
}
