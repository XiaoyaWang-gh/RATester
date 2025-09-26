package hugofs

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestType_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := dirEntry{}
	if got := d.Type(); got != fs.FileMode(0) {
		t.Errorf("dirEntry.Type() = %v, want %v", got, fs.FileMode(0))
	}
}
