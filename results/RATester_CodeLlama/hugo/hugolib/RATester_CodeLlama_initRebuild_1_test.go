package hugolib

import (
	"fmt"
	"testing"
)

func TestInitRebuild_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	t.Parallel()
	var h HugoSites
	var config *BuildCfg
	if err := h.initRebuild(config); err != nil {
		t.Fatal(err)
	}
}
