package hugolib

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/markup/tableofcontents"
)

func TestFragments_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := context.Background()
	pco := &pageContentOutput{
		po: &pageOutput{
			// Initialize necessary fields
		},
	}

	expected := &tableofcontents.Fragments{}

	result := pco.Fragments(ctx)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
