package test

import (
	"fmt"
	"testing"
)

func TestAssetNames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	names := AssetNames()
	if len(names) == 0 {
		t.Error("AssetNames() should not return an empty slice")
	}
}
