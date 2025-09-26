package hugofs

import (
	"fmt"
	"testing"

	radix "github.com/armon/go-radix"
)

func TestFilter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := RootMappingFs{
		id:            "test",
		rootMapToReal: radix.New(),
	}

	fs.rootMapToReal.Insert("test1", []RootMapping{
		{From: "test1"},
		{From: "test2"},
	})

	fs.rootMapToReal.Insert("test2", []RootMapping{
		{From: "test3"},
		{From: "test4"},
	})

	filteredFs := fs.Filter(func(m RootMapping) bool {
		return m.From != "test1"
	})

	filteredFs.rootMapToReal.Walk(func(b string, v any) bool {
		rms := v.([]RootMapping)
		for _, rm := range rms {
			if rm.From == "test1" {
				t.Errorf("Expected RootMapping with From 'test1' to be filtered out")
			}
		}
		return false
	})
}
