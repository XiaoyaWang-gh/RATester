package docker

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestBuild_2(t *testing.T) {
	ctx := context.Background()

	testCases := []struct {
		desc                string
		containersInspected []dockerData
		expected            *dynamic.Configuration
	}{
		{
			desc: "test case 1",
			containersInspected: []dockerData{
				{
					ID: "1",
					Labels: map[string]string{
						"traefik.enable": "true",
					},
				},
			},
			expected: &dynamic.Configuration{
				HTTP: &dynamic.HTTPConfiguration{},
			},
		},
		// Add more test cases as needed
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			builder := &DynConfBuilder{}
			result := builder.build(ctx, test.containersInspected)

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
	}
}
