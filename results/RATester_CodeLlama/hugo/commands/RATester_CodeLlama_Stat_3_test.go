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

	fs := &countingStatFs{Fs: afero.NewMemMapFs()}
	name := "test.txt"
	_, err := fs.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	_, err = fs.Stat(name)
	if err != nil {
		t.Fatal(err)
	}
	if fs.statCounter != 1 {
		t.Errorf("statCounter = %d, want 1", fs.statCounter)
	}
}
