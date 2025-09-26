package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/fsnotify/fsnotify"
	"github.com/gohugoio/hugo/hugolib/filesystems"
)

func TestpartitionDynamicEvents_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	sourceFs := &filesystems.SourceFilesystems{
		Content: &filesystems.SourceFilesystem{},
		Data:    &filesystems.SourceFilesystem{},
		I18n:    &filesystems.SourceFilesystem{},
		Layouts: &filesystems.SourceFilesystem{},
		Static: map[string]*filesystems.SourceFilesystem{
			"": &filesystems.SourceFilesystem{},
		},
	}

	events := []fsnotify.Event{
		{Name: "static/file1.txt", Op: fsnotify.Create},
		{Name: "content/file2.txt", Op: fsnotify.Create},
		{Name: "data/file3.txt", Op: fsnotify.Create},
		{Name: "i18n/file4.txt", Op: fsnotify.Create},
		{Name: "layouts/file5.txt", Op: fsnotify.Create},
	}

	expectedContentEvents := []fsnotify.Event{
		{Name: "content/file2.txt", Op: fsnotify.Create},
		{Name: "data/file3.txt", Op: fsnotify.Create},
		{Name: "i18n/file4.txt", Op: fsnotify.Create},
		{Name: "layouts/file5.txt", Op: fsnotify.Create},
	}

	expectedAssetEvents := []fsnotify.Event{
		{Name: "static/file1.txt", Op: fsnotify.Create},
	}

	de := partitionDynamicEvents(sourceFs, events)

	if !reflect.DeepEqual(de.ContentEvents, expectedContentEvents) {
		t.Errorf("Expected ContentEvents to be %v, but got %v", expectedContentEvents, de.ContentEvents)
	}

	if !reflect.DeepEqual(de.AssetEvents, expectedAssetEvents) {
		t.Errorf("Expected AssetEvents to be %v, but got %v", expectedAssetEvents, de.AssetEvents)
	}
}
