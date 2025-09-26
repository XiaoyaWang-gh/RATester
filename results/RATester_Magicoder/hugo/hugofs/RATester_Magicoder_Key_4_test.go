package hugofs

import (
	"fmt"
	"testing"
)

func TestKey_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &RootMappingFs{
		id: "testID",
	}

	expected := "testID"
	actual := fs.Key()

	if actual != expected {
		t.Errorf("Expected Key() to return %s, but got %s", expected, actual)
	}
}
