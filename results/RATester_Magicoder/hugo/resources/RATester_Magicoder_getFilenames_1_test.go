package resources

import (
	"fmt"
	"testing"
)

func TestgetFilenames_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	cache := &ResourceCache{}
	key := "testKey"
	metaFilename, contentFilename := cache.getFilenames(key)

	if metaFilename != key+".json" {
		t.Errorf("Expected meta filename to be %s, but got %s", key+".json", metaFilename)
	}

	if contentFilename != key+".content" {
		t.Errorf("Expected content filename to be %s, but got %s", key+".content", contentFilename)
	}
}
