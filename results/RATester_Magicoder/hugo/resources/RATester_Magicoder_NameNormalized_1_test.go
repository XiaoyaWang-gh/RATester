package resources

import (
	"fmt"
	"testing"
)

func TestNameNormalized_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		sd: ResourceSourceDescriptor{
			NameNormalized: "test",
		},
	}

	expected := "test"
	result := l.NameNormalized()

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
