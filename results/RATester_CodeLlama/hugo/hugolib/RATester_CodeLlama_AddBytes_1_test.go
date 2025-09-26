package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/parser/pageparser"
)

func TestAddBytes_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var item pageparser.Item
	var p *contentParseInfo
	// When
	p.AddBytes(item)
	// Then
	// TODO: Add test cases.
}
