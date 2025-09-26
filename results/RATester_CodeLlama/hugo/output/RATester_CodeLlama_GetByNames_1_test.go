package output

import (
	"fmt"
	"testing"
)

func TestGetByNames_1(t *testing.T) {
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

	names := []string{"html", "rss"}

	types, err := formats.GetByNames(names...)
	if err != nil {
		t.Fatal(err)
	}

	if len(types) != len(names) {
		t.Fatalf("Expected %d types, got %d", len(names), len(types))
	}

	for i, name := range names {
		if types[i].Name != name {
			t.Errorf("Expected type %q, got %q", name, types[i].Name)
		}
	}
}
