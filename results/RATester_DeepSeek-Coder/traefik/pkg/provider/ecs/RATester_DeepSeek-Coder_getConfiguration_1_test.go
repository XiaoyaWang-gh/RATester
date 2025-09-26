package ecs

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetConfiguration_1(t *testing.T) {
	testCases := []struct {
		desc     string
		instance ecsInstance
		wantConf configuration
		wantErr  bool
	}{
		{
			desc: "should return error when label decoding fails",
			instance: ecsInstance{
				Labels: map[string]string{
					"traefik.ecs.enable": "invalid",
				},
			},
			wantErr: true,
		},
		{
			desc: "should return configuration when label decoding succeeds",
			instance: ecsInstance{
				Labels: map[string]string{
					"traefik.ecs.enable": "true",
				},
			},
			wantConf: configuration{
				Enable: true,
			},
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			provider := &Provider{
				ExposedByDefault: true,
			}

			conf, err := provider.getConfiguration(test.instance)
			if (err != nil) != test.wantErr {
				t.Errorf("getConfiguration() error = %v, wantErr %v", err, test.wantErr)
				return
			}

			if !reflect.DeepEqual(conf, test.wantConf) {
				t.Errorf("getConfiguration() = %v, want %v", conf, test.wantConf)
			}
		})
	}
}
