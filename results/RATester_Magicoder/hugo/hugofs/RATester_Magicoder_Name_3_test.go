package hugofs

import (
	"fmt"
	"testing"
)

func TestName_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := &rootMappingDir{
		name: "test_dir",
	}

	expected := "test_dir"
	actual := dir.Name()

	if actual != expected {
		t.Errorf("Expected %s, but got %s", expected, actual)
	}
}
