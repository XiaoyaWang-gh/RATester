package consulcatalog

import (
	"fmt"
	"testing"
)

func TestInit_21(t *testing.T) {
	testCases := []struct {
		desc     string
		provider *Provider
		expErr   bool
	}{
		{
			desc: "should return error when default rule is not valid",
			provider: &Provider{
				Configuration: Configuration{
					DefaultRule: "invalid",
				},
			},
			expErr: true,
		},
		{
			desc: "should not return error when default rule is valid",
			provider: &Provider{
				Configuration: Configuration{
					DefaultRule: "Host(`{{ normalize .Provider.Service }}`)",
				},
			},
			expErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := test.provider.Init()
			if (err != nil) != test.expErr {
				t.Errorf("Init() error = %v, wantErr %v", err, test.expErr)
			}
		})
	}
}
