package hugofs

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/paths"
)

func TestStat_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	dir := &rootMappingDir{
		name: "testDir",
		meta: &FileMeta{
			PathInfo: &paths.Path{},
		},
	}

	_, err := dir.Stat()
	if err != errIsDir {
		t.Errorf("Expected error %v, but got %v", errIsDir, err)
	}
}
