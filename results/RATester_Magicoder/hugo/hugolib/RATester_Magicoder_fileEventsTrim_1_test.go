package hugolib

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestfileEventsTrim_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}
	events := []fsnotify.Event{
		{Name: "file1"},
		{Name: "file2"},
		{Name: "file1"},
		{Name: "file3"},
		{Name: "file2"},
	}
	expected := []fsnotify.Event{
		{Name: "file1"},
		{Name: "file2"},
		{Name: "file3"},
	}
	result := h.fileEventsTrim(events)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
