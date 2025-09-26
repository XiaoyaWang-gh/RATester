package paths

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestLang_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &Path{
		identifiers: []types.LowHigh[string]{
			{Low: 1, High: 1},
			{Low: 2, High: 2},
			{Low: 3, High: 3},
		},
	}

	if got := p.Lang(); got != "1" {
		t.Errorf("Lang() = %v, want %v", got, "1")
	}
}
