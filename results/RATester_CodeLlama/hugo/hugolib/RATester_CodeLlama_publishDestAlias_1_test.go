package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestPublishDestAlias_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	s := &Site{}
	allowRoot := true
	path := "path"
	permalink := "permalink"
	outputFormat := output.Format{Name: "outputFormat"}
	p := &pageState{}

	// When
	err := s.publishDestAlias(allowRoot, path, permalink, outputFormat, p)

	// Then
	if err != nil {
		t.Errorf("publishDestAlias() error = %v", err)
		return
	}
}
