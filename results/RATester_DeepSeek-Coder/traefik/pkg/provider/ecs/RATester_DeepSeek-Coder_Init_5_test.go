package ecs

import (
	"fmt"
	"testing"
)

func TestInit_5(t *testing.T) {
	testCases := []struct {
		desc          string
		givenProvider *Provider
		wantErr       bool
	}{
		{
			desc: "should return error when default rule is invalid",
			givenProvider: &Provider{
				DefaultRule: "{{ invalid_rule }}",
			},
			wantErr: true,
		},
		{
			desc: "should not return error when default rule is valid",
			givenProvider: &Provider{
				DefaultRule: "PathPrefix(`/api`)",
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

			err := tC.givenProvider.Init()
			if (err != nil) != tC.wantErr {
				t.Errorf("Init() error = %v, wantErr %v", err, tC.wantErr)
			}
		})
	}
}
