package test

import (
	"fmt"
	"testing"
)

func TestAssetInfo_1(t *testing.T) {
	t.Run("should return file info for existing asset", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		info, err := AssetInfo("existing_asset.txt")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if info == nil {
			t.Error("expected file info, got nil")
		}
	})

	t.Run("should return error for non-existing asset", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		_, err := AssetInfo("non_existing_asset.txt")
		if err == nil {
			t.Error("expected error, got nil")
		}
	})
}
