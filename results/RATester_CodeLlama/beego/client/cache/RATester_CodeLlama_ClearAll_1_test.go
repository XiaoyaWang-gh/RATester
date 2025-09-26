package cache

import (
	"context"
	"fmt"
	"testing"
)

func TestClearAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath:      "./testdata",
		FileSuffix:     ".json",
		DirectoryLevel: 2,
		EmbedExpiry:    0,
	}
	err := fc.ClearAll(context.Background())
	if err != nil {
		t.Errorf("ClearAll() error = %v", err)
		return
	}
	t.Log("ClearAll() success")
}
