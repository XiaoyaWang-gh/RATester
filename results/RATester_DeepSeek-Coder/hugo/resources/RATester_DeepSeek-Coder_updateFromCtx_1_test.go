package resources

import (
	"fmt"
	"reflect"
	"testing"
)

func TestUpdateFromCtx_1(t *testing.T) {
	type args struct {
		ctx *ResourceTransformationCtx
	}
	tests := []struct {
		name string
		u    *transformationUpdate
		args args
		want *transformationUpdate
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

			u := &transformationUpdate{}
			u.updateFromCtx(tt.args.ctx)
			if !reflect.DeepEqual(u, tt.want) {
				t.Errorf("updateFromCtx() = %v, want %v", u, tt.want)
			}
		})
	}
}
