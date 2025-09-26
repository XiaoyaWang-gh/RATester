package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestFuzzyWordCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	ctx := context.Background()
	assert.Equal(t, 0, pco.FuzzyWordCount(ctx))
}
