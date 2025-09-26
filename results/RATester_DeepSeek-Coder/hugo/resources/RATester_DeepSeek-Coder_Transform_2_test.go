package resources

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/herrors"
)

func TestTransform_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	ctx := &ResourceTransformationCtx{
		Ctx: context.Background(),
	}

	transformer := transformerNotAvailable{}
	err := transformer.Transform(ctx)

	if err != herrors.ErrFeatureNotAvailable {
		t.Errorf("Expected error %v, got %v", herrors.ErrFeatureNotAvailable, err)
	}
}
