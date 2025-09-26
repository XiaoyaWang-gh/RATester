package ecs

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/stretchr/testify/require"
	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuildServiceConfiguration_1(t *testing.T) {
	testCases := []struct {
		desc     string
		instance ecsInstance
		want     *dynamic.HTTPConfiguration
	}{
		{
			desc: "should build service configuration",
			instance: ecsInstance{
				Name: "test",
				containerDefinition: &ecs.ContainerDefinition{
					DockerLabels: map[string]*string{
						"traefik.http.services.test.loadbalancer.server.port": aws.String("80"),
					},
				},
			},
			want: &dynamic.HTTPConfiguration{
				Services: map[string]*dynamic.Service{
					"test": {
						LoadBalancer: &dynamic.ServersLoadBalancer{
							Servers: []dynamic.Server{
								{
									URL: "http://127.0.0.1:80",
								},
							},
						},
					},
				},
			},
		},
	}

	for _, test := range testCases {
		test := test
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			t.Parallel()

			p := &Provider{}
			configuration := &dynamic.HTTPConfiguration{}
			err := p.buildServiceConfiguration(context.Background(), test.instance, configuration)
			require.NoError(t, err)
			require.Equal(t, test.want, configuration)
		})
	}
}
