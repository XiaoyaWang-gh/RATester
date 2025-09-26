package schema

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestContainsAlias_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	infos := []*structInfo{
		{
			fields: []*fieldInfo{
				{
					alias: "A",
				},
				{
					alias: "B",
				},
			},
		},
		{
			fields: []*fieldInfo{
				{
					alias: "C",
				},
				{
					alias: "D",
				},
			},
		},
	}
	alias := "B"

	// when
	result := containsAlias(infos, alias)

	// then
	assert.True(t, result)
}
