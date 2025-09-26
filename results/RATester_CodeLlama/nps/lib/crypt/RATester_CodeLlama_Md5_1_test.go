package crypt

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestMd5_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	assert := assert.New(t)
	assert.Equal("900150983cd24fb0d6963f7d28e17f72", Md5("123456"))
}
