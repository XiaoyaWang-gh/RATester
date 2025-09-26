package assets

import (
	"fmt"
	"testing"
	"testing/fstest"
)

func TestRegister_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &fstest.MapFS{
		"static": &fstest.MapFile{
			Data: []byte("static"),
		},
	}
	Register(fs)
	if content != fs {
		t.Errorf("Register() = %v, want %v", content, fs)
	}
}
