package hugo

import (
	"fmt"
	"html/template"
	"testing"
)

func TestGenerator_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	info := HugoInfo{
		CommitHash:  "abcdefg",
		BuildDate:   "2022-01-01",
		Environment: "production",
		GoVersion:   "1.17",
	}

	expected := `<meta name="generator" content="Hugo 0.80.0">`
	result := info.Generator()

	if result != template.HTML(expected) {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
