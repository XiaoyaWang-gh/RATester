package related

import (
	"fmt"
	"testing"
)

func TestGetIndexCfg_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	name := "test"
	cfg := Config{
		Indices: IndicesConfig{
			{
				Name: "test",
			},
		},
	}
	idx := &InvertedIndex{
		cfg: cfg,
	}

	// Act
	result, ok := idx.getIndexCfg(name)

	// Assert
	if !ok {
		t.Errorf("Expected ok to be true, but was %v", ok)
	}
	if result.Name != name {
		t.Errorf("Expected result.Name to be %q, but was %q", name, result.Name)
	}
}
