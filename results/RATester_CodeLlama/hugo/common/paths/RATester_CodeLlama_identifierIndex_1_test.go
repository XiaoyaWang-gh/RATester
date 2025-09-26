package paths

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestIdentifierIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		identifiers: []types.LowHigh[string]{
			{Low: 1, High: 2},
			{Low: 3, High: 4},
			{Low: 5, High: 6},
		},
	}

	for i := -1; i < 6; i++ {
		if got := p.identifierIndex(i); got != i {
			t.Errorf("identifierIndex(%d) = %d, want %d", i, got, i)
		}
	}
}
