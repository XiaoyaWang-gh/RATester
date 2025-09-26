package hugofs

import (
	"fmt"
	"testing"
)

func TestNewFileMeta_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fm := NewFileMeta()
	if fm == nil {
		t.Error("NewFileMeta() should not return nil")
	}
	if fm.PathInfo != nil {
		t.Error("PathInfo should be nil")
	}
	if fm.Name != "" {
		t.Error("Name should be empty")
	}
	if fm.Filename != "" {
		t.Error("Filename should be empty")
	}
	if fm.BaseDir != "" {
		t.Error("BaseDir should be empty")
	}
	if fm.SourceRoot != "" {
		t.Error("SourceRoot should be empty")
	}
	if fm.Module != "" {
		t.Error("Module should be empty")
	}
	if fm.ModuleOrdinal != 0 {
		t.Error("ModuleOrdinal should be 0")
	}
	if fm.Component != "" {
		t.Error("Component should be empty")
	}
	if fm.Weight != 0 {
		t.Error("Weight should be 0")
	}
	if fm.IsProject != false {
		t.Error("IsProject should be false")
	}
	if fm.Watch != false {
		t.Error("Watch should be false")
	}
	if fm.Lang != "" {
		t.Error("Lang should be empty")
	}
	if fm.LangIndex != 0 {
		t.Error("LangIndex should be 0")
	}
	if fm.OpenFunc != nil {
		t.Error("OpenFunc should be nil")
	}
	if fm.JoinStatFunc != nil {
		t.Error("JoinStatFunc should be nil")
	}
	if fm.InclusionFilter != nil {
		t.Error("InclusionFilter should be nil")
	}
	if fm.Rename != nil {
		t.Error("Rename should be nil")
	}
}
