package hugofs

import (
	"fmt"
	"io/fs"
	"testing"
)

func TestMkdir_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	name := "test"
	perm := fs.FileMode(0755)
	fs := noOpFs{}
	err := fs.Mkdir(name, perm)
	if err != nil {
		t.Errorf("Mkdir() error = %v, want nil", err)
		return
	}
}
