package hugofs

import (
	"fmt"
	"testing"

	"github.com/spf13/afero"
)

func TestNewCreateCountingFs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fs := afero.NewMemMapFs()
	fs = NewCreateCountingFs(fs)
	if _, ok := fs.(*createCountingFs); !ok {
		t.Errorf("NewCreateCountingFs() = %T, want %T", fs, &createCountingFs{})
	}
}
