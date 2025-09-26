package types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeepCopyInto_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := &Domain{
		Main: "foo.com",
		SANs: []string{"bar.com", "baz.com"},
	}
	out := &Domain{}
	in.DeepCopyInto(out)
	assert.Equal(t, in, out)
	assert.NotSame(t, in, out)
	assert.NotSame(t, in.SANs, out.SANs)
}
