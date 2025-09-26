package pagemeta

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestValueAsOpenReadSeekCloser_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	s := Source{
		Value: "foo",
	}

	// when
	actual := s.ValueAsOpenReadSeekCloser()

	// then
	assert.NotNil(t, actual)
}
