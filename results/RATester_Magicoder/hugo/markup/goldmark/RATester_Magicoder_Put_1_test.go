package goldmark

import (
	"fmt"
	"testing"

	"github.com/yuin/goldmark/util"
)

func TestPut_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ids := &idFactory{
		idType: "test",
		vals:   make(map[string]struct{}),
	}

	value := []byte("test_value")
	ids.Put(value)

	if _, ok := ids.vals[util.BytesToReadOnlyString(value)]; !ok {
		t.Errorf("Expected value to be in the map, but it was not.")
	}
}
