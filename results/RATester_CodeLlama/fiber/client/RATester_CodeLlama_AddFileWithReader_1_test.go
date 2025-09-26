package client

import (
	"fmt"
	"io/ioutil"
	"strings"
	"testing"

	"gotest.tools/assert"
)

func TestAddFileWithReader_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// given
	r := &Request{}
	name := "test.txt"
	reader := ioutil.NopCloser(strings.NewReader("test"))
	// when
	r.AddFileWithReader(name, reader)
	// then
	assert.Equal(t, 1, len(r.files))
	assert.Equal(t, name, r.files[0].name)
	assert.Equal(t, reader, r.files[0].reader)
}
