package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
)

func TestGetLeaf_1(t *testing.T) {
	type fields struct {
		root []string
		leaf keyPair
	}
	tests := []struct {
		name   string
		fields fields
		want   traefiktls.Certificate
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

			c := &connectCert{
				root: tt.fields.root,
				leaf: tt.fields.leaf,
			}
			if got := c.getLeaf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectCert.getLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
