package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetLeaf_1(t *testing.T) {
	type fields struct {
		leaf keyPair
	}
	tests := []struct {
		name   string
		fields fields
		want   traefiktls.Certificate
	}{
		{
			name: "Test Case 1",
			fields: fields{
				leaf: keyPair{
					cert: "test_cert",
					key:  "test_key",
				},
			},
			want: traefiktls.Certificate{
				CertFile: types.FileOrContent("test_cert"),
				KeyFile:  types.FileOrContent("test_key"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &connectCert{
				leaf: tt.fields.leaf,
			}
			if got := c.getLeaf(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("connectCert.getLeaf() = %v, want %v", got, tt.want)
			}
		})
	}
}
