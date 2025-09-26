package nomad

import (
	"fmt"
	"testing"

	"github.com/hashicorp/nomad/api"
)

func TestInit_14(t *testing.T) {
	testCases := []struct {
		desc     string
		provider *Provider
		wantErr  bool
	}{
		{
			desc: "should return error when namespace is wildcard",
			provider: &Provider{
				namespace: api.AllNamespacesNamespace,
			},
			wantErr: true,
		},
		{
			desc: "should not return error when namespace is not wildcard",
			provider: &Provider{
				namespace: "not-wildcard",
			},
			wantErr: false,
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tC.provider.Init()
			if (err != nil) != tC.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tC.wantErr)
			}
		})
	}
}
