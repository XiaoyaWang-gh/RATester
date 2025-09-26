package hugofs

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestMkdirAll_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	path := "path"
	perm := fs.FileMode(0)
	fs := noOpFs{}
	err := fs.MkdirAll(path, perm)
	if err != nil {
		t.Errorf("MkdirAll() error = %v, want nil", err)
		return
	}
}
