package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetRelatedDocsHandler_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	s := &Site{}
	assert.NotNil(t, s.GetRelatedDocsHandler())
}
