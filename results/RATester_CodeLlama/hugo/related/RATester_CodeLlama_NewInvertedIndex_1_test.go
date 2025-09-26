package related

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewInvertedIndex_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	cfg := Config{
		Threshold: 100,
		Indices: IndicesConfig{
			{
				Name: "index1",
			},
			{
				Name: "index2",
			},
		},
	}

	// when
	idx := NewInvertedIndex(cfg)

	// then
	assert.NotNil(t, idx)
	assert.Equal(t, 2, len(idx.index))
	assert.Equal(t, 100, idx.minWeight)
	assert.Equal(t, 100, idx.maxWeight)
}
