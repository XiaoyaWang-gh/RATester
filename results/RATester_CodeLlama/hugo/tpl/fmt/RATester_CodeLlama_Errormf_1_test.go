package fmt

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestErrormf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ns := &Namespace{}
	m := "test"
	format := "test %s"
	args := []any{"test"}
	assert.Equal(t, "test test", ns.Errormf(m, format, args...))
}
