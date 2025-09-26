package paths

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestIdentifiers_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		s: "a/b/c/d/e/f",
		identifiers: []types.LowHigh[string]{
			{Low: 1, High: 2},
			{Low: 3, High: 4},
			{Low: 5, High: 6},
		},
	}

	expected := []string{"b", "c", "d"}
	if !reflect.DeepEqual(p.Identifiers(), expected) {
		t.Errorf("Identifiers() = %v, want %v", p.Identifiers(), expected)
	}
}
