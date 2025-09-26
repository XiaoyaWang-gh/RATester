package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/output"
)

func TestWriteDestAlias_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	s := &Site{}
	path := "path"
	permalink := "permalink"
	outputFormat := output.Format{}
	p := &pageState{}

	// When
	err := s.writeDestAlias(path, permalink, outputFormat, p)

	// Then
	if err != nil {
		t.Errorf("writeDestAlias() error = %v", err)
		return
	}
}
