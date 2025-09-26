package pagemeta

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsDateKey_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	key := "date"
	f := FrontMatterHandler{}
	f.allDateKeys = map[string]bool{
		"date": true,
	}

	// when
	result := f.IsDateKey(key)

	// then
	assert.True(t, result)
}
