package openapi3

import (
	"fmt"
	"reflect"
	"testing"

	kopenapi3 "github.com/getkin/kin-openapi/openapi3"
	"github.com/gohugoio/hugo/identity"
)

func TestGetIdentityGroup_2(t *testing.T) {
	type fields struct {
		T             *kopenapi3.T
		identityGroup identity.Identity
	}
	tests := []struct {
		name   string
		fields fields
		want   identity.Identity
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

			o := &OpenAPIDocument{
				T:             tt.fields.T,
				identityGroup: tt.fields.identityGroup,
			}
			if got := o.GetIdentityGroup(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("OpenAPIDocument.GetIdentityGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
