package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/maps"
)

func TestnewPageContentOutput_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	po := &pageOutput{
		p: &pageState{
			pid: 1,
		},
	}
	expected := &pageContentOutput{
		po:           po,
		renderHooks:  &renderHooks{},
		otherOutputs: maps.NewCache[uint64, *pageContentOutput](),
	}
	actual, err := newPageContentOutput(po)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expected: %v, but got: %v", expected, actual)
	}
}
