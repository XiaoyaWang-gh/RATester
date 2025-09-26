package resources

import (
	"fmt"
	"testing"
)

func TestUpdateFromCtx_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &ResourceTransformationCtx{
		InPath: "path/to/resource",
	}
	u := &transformationUpdate{}
	u.updateFromCtx(ctx)
	if u.targetPath != ctx.InPath {
		t.Errorf("Expected targetPath to be %q but was %q", ctx.InPath, u.targetPath)
	}
}
