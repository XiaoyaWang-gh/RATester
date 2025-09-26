package hugofs

import (
	"fmt"
	"testing"
)

func TestChown_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := &noOpFs{}
	name := "test"
	uid := 1
	gid := 1
	err := fs.Chown(name, uid, gid)
	if err != errNoOp {
		t.Errorf("Chown() = %v, want %v", err, errNoOp)
	}
}
