package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/spf13/afero"
)

func TestNewHasBytesReceiver_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	delegate := afero.NewMemMapFs()
	shouldCheck := func(name string) bool {
		return true
	}
	hasBytesCallback := func(name string, match []byte) {
		// do nothing
	}
	patterns := [][]byte{[]byte("pattern1"), []byte("pattern2")}

	// when
	fs := NewHasBytesReceiver(delegate, shouldCheck, hasBytesCallback, patterns...)

	// then
	assert.NotNil(t, fs)
}
