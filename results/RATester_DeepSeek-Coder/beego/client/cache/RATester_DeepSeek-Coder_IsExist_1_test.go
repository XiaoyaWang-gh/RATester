package cache

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestIsExist_1(t *testing.T) {
	fc := &FileCache{
		CachePath: "/tmp",
	}

	t.Run("ExistingFile", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		file, err := ioutil.TempFile(fc.CachePath, "")
		if err != nil {
			t.Fatalf("Failed to create temp file: %v", err)
		}
		defer os.Remove(file.Name())

		ctx := context.Background()
		exist, err := fc.IsExist(ctx, file.Name())
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if !exist {
			t.Errorf("Expected file to exist, but it does not")
		}
	})

	t.Run("NonExistingFile", func(t *testing.T) {

		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Recovered in main", r)
			}
		}()

		ctx := context.Background()
		exist, err := fc.IsExist(ctx, "non-existing-file")
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if exist {
			t.Errorf("Expected file to not exist, but it does")
		}
	})
}
