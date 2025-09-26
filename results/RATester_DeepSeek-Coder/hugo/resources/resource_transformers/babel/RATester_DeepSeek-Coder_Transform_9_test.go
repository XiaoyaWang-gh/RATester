package babel

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources"
)

func TestTransform_9(t *testing.T) {
	type args struct {
		ctx *resources.ResourceTransformationCtx
	}
	tests := []struct {
		name    string
		t       *babelTransformation
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.t.Transform(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("babelTransformation.Transform() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
