package consulcatalog

import (
	"fmt"
	"reflect"
	"testing"

	traefiktls "github.com/traefik/traefik/v3/pkg/tls"
	"github.com/traefik/traefik/v3/pkg/types"
)

func TestGetLeaf_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &connectCert{
		leaf: keyPair{
			cert: "cert",
			key:  "key",
		},
	}
	want := traefiktls.Certificate{
		CertFile: types.FileOrContent("cert"),
		KeyFile:  types.FileOrContent("key"),
	}
	if got := c.getLeaf(); !reflect.DeepEqual(got, want) {
		t.Errorf("getLeaf() = %v, want %v", got, want)
	}
}
