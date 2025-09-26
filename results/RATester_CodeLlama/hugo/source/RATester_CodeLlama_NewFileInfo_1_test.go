package source

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestNewFileInfo_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var fi hugofs.FileMetaInfo
	got := NewFileInfo(fi)
	if got == nil {
		t.Errorf("NewFileInfo() = %v, want %v", got, &File{})
	}
}
