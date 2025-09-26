package kv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSetDefaults_13(t *testing.T) {
	testCases := []struct {
		desc string
		have *Provider
		want *Provider
	}{
		{
			desc: "should set default root key",
			have: &Provider{},
			want: &Provider{RootKey: "traefik"},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tC.have.SetDefaults()

			if !reflect.DeepEqual(tC.have, tC.want) {
				t.Errorf("have %v, want %v", tC.have, tC.want)
			}
		})
	}
}
