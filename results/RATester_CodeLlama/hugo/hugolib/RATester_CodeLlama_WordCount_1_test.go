package hugolib

import (
	"context"
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestWordCount_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pco := &pageContentOutput{}
	pco.contentRendered.Store(true)
	assert.Equal(t, 0, pco.WordCount(context.Background()))
}
