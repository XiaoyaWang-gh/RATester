package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAddOutPathIdentifier_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &ResourceTransformationCtx{
		InPath: "path/to/file.txt",
	}

	ctx.AddOutPathIdentifier("identifier")

	assert.Equal(t, "path/to/file.identifier.txt", ctx.OutPath)
}
