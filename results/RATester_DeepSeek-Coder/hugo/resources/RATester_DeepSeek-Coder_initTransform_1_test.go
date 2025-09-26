package resources

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/resources/resource"
)

func TestInitTransform_1(t *testing.T) {
	type fields struct {
		commonResource          commonResource
		resourceTransformations *resourceTransformations
		resourceAdapterInner    *resourceAdapterInner
		metaProvider            resource.ResourceMetaProvider
	}
	type args struct {
		publish    bool
		setContent bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			r := &resourceAdapter{
				commonResource:          tt.fields.commonResource,
				resourceTransformations: tt.fields.resourceTransformations,
				resourceAdapterInner:    tt.fields.resourceAdapterInner,
				metaProvider:            tt.fields.metaProvider,
			}
			r.initTransform(tt.args.publish, tt.args.setContent)
		})
	}
}
