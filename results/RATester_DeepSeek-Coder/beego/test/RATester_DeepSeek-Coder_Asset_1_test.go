package test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestAsset_1(t *testing.T) {
	t.Run("ExistingAsset", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "existing_asset.txt"
		expectedContent := []byte("This is an existing asset.")
		_bindata[name] = func() (*asset, error) {
			return &asset{
				bytes: expectedContent,
				info:  nil,
			}, nil
		}

		content, err := Asset(name)
		if err != nil {
			t.Errorf("Expected no error, got %v", err)
		}
		if !bytes.Equal(content, expectedContent) {
			t.Errorf("Expected content %v, got %v", expectedContent, content)
		}
	})

	t.Run("NonExistingAsset", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "non_existing_asset.txt"
		expectedErr := fmt.Errorf("Asset %s not found", name)

		_, err := Asset(name)
		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}
	})

	t.Run("AssetReadError", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		name := "asset_read_error.txt"
		expectedErr := fmt.Errorf("Asset %s can't read by error: read error", name)
		_bindata[name] = func() (*asset, error) {
			return nil, expectedErr
		}

		_, err := Asset(name)
		if err == nil || err.Error() != expectedErr.Error() {
			t.Errorf("Expected error %v, got %v", expectedErr, err)
		}
	})
}
