package commands

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestStat_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &countingStatFs{
		Fs:          afero.NewMemMapFs(),
		statCounter: 0,
	}

	_, err := fs.Stat("test.txt")
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if fs.statCounter != 1 {
		t.Errorf("Expected statCounter to be 1, got %d", fs.statCounter)
	}

	_, err = fs.Stat("non-existent.txt")
	if err == nil {
		t.Error("Expected error, got nil")
	}

	if fs.statCounter != 1 {
		t.Errorf("Expected statCounter to be 1, got %d", fs.statCounter)
	}
}
