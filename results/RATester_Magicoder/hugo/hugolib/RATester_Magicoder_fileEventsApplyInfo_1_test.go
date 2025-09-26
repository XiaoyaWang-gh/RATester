package hugolib

import (
	"fmt"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestfileEventsApplyInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	h := &HugoSites{}

	testEvents := []fsnotify.Event{
		{Name: "test1", Op: fsnotify.Remove},
		{Name: "test2", Op: fsnotify.Rename},
		{Name: "test3", Op: fsnotify.Create},
	}

	expectedInfos := []fileEventInfo{
		{Event: fsnotify.Event{Name: "test1", Op: fsnotify.Remove}, added: false, removed: true, isChangedDir: false},
		{Event: fsnotify.Event{Name: "test2", Op: fsnotify.Rename}, added: false, removed: true, isChangedDir: false},
		{Event: fsnotify.Event{Name: "test3", Op: fsnotify.Create}, added: true, removed: false, isChangedDir: false},
	}

	infos := h.fileEventsApplyInfo(testEvents)

	if len(infos) != len(expectedInfos) {
		t.Errorf("Expected %d infos, got %d", len(expectedInfos), len(infos))
	}

	for i, info := range infos {
		if info.Event.Name != expectedInfos[i].Event.Name || info.Event.Op != expectedInfos[i].Event.Op || info.added != expectedInfos[i].added || info.removed != expectedInfos[i].removed || info.isChangedDir != expectedInfos[i].isChangedDir {
			t.Errorf("Expected info %+v, got %+v", expectedInfos[i], info)
		}
	}
}
