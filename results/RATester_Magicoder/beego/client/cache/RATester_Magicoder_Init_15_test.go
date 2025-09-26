package cache

import (
	"fmt"
	"os"
	"testing"
)

func TestInit_15(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fc := &FileCache{
		CachePath: "/tmp/test",
	}

	err := fc.Init()
	if err != nil {
		t.Errorf("Init() failed, error: %v", err)
	}

	_, err = os.Stat(fc.CachePath)
	if os.IsNotExist(err) {
		t.Errorf("Init() failed, directory not created: %v", err)
	}

	os.RemoveAll(fc.CachePath)
}
