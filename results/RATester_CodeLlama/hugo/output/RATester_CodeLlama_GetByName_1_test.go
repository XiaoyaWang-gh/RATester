package output

import (
	"fmt"
	"testing"
)

func TestGetByName_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	formats := Formats{
		{
			Name: "html",
		},
		{
			Name: "rss",
		},
	}

	f, found := formats.GetByName("html")
	if !found {
		t.Error("expected found to be true")
	}
	if f.Name != "html" {
		t.Errorf("expected f.Name to be %q but was %q", "html", f.Name)
	}

	f, found = formats.GetByName("rss")
	if !found {
		t.Error("expected found to be true")
	}
	if f.Name != "rss" {
		t.Errorf("expected f.Name to be %q but was %q", "rss", f.Name)
	}

	f, found = formats.GetByName("json")
	if found {
		t.Error("expected found to be false")
	}
	if f.Name != "" {
		t.Errorf("expected f.Name to be %q but was %q", "", f.Name)
	}
}
