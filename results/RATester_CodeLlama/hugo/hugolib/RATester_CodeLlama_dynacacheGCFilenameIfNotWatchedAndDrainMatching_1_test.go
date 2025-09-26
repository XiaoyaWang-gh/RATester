package hugolib

import (
	"fmt"
	"testing"
)

func TestDynacacheGCFilenameIfNotWatchedAndDrainMatching_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var h HugoSites
	var filename string
	h.dynacacheGCFilenameIfNotWatchedAndDrainMatching(filename)
}
