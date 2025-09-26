package commands

import (
	"fmt"
	"testing"
)

func TestbuildSites_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	hb := &hugoBuilder{}
	err := hb.buildSites(false)
	if err != nil {
		t.Errorf("buildSites() error = %v, wantErr %v", err, false)
		return
	}
}
