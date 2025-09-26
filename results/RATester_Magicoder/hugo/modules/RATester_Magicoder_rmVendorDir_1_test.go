package modules

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
)

func TestrmVendorDir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	client := &Client{
		fs: afero.NewMemMapFs(),
	}

	vendorDir := "vendor"
	modulestxt := filepath.Join(vendorDir, vendorModulesFilename)

	// Test case 1: Vendor directory does not exist
	err := client.rmVendorDir(vendorDir)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}

	// Test case 2: Vendor directory exists but modules.txt does not exist
	err = client.fs.MkdirAll(vendorDir, 0755)
	if err != nil {
		t.Errorf("Failed to create vendor directory: %v", err)
	}

	err = client.rmVendorDir(vendorDir)
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}

	// Test case 3: Vendor directory and modules.txt exist
	err = client.fs.MkdirAll(vendorDir, 0755)
	if err != nil {
		t.Errorf("Failed to create vendor directory: %v", err)
	}

	_, err = client.fs.Create(modulestxt)
	if err != nil {
		t.Errorf("Failed to create modules.txt: %v", err)
	}

	err = client.rmVendorDir(vendorDir)
	if err != nil {
		t.Errorf("Expected no error, but got %v", err)
	}
}
