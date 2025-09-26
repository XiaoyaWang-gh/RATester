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

	f := &rootMappingDir{}
	f.name = "test"
	if f.Name() != "test" {
		t.Errorf("Expected %s, got %s", "test", f.Name())
	}
}
