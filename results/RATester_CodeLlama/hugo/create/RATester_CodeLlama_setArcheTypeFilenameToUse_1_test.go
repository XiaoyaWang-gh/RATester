package create

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/hugofs"
)

func TestSetArcheTypeFilenameToUse_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var b *contentBuilder
	var ext string
	var pathsToCheck []string

	if b.kind != "" {
		pathsToCheck = append(pathsToCheck, b.kind+ext)
	}

	pathsToCheck = append(pathsToCheck, "default"+ext)

	for _, p := range pathsToCheck {
		fi, err := b.archeTypeFs.Stat(p)
		if err == nil {
			b.archetypeFi = fi.(hugofs.FileMetaInfo)
			b.isDir = fi.IsDir()
			return
		}
	}
}
