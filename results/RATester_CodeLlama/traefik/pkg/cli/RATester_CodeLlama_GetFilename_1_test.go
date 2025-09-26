package cli

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestGetFilename_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := &FileLoader{
		ConfigFileFlag: "test.yml",
	}
	assert.Equal(t, "test.yml", f.GetFilename())
}
