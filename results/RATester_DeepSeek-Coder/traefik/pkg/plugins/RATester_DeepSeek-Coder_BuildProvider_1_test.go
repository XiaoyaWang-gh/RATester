package plugins

import (
	"fmt"
	"testing"
)

func TestBuildProvider_1(t *testing.T) {
	type testCase struct {
		name          string
		builder       Builder
		pName         string
		config        map[string]interface{}
		expectedError bool
	}

	testCases := []testCase{
		{
			name: "success",
			builder: Builder{
				providerBuilders: map[string]providerBuilder{
					"test": {
						Import: "test",
					},
				},
			},
			pName: "test",
			config: map[string]interface{}{
				"test": "config",
			},
			expectedError: false,
		},
		{
			name: "no plugin definition in the static configuration",
			builder: Builder{
				providerBuilders: map[string]providerBuilder{
					"test": {
						Import: "test",
					},
				},
			},
			pName: "unknown",
			config: map[string]interface{}{
				"test": "config",
			},
			expectedError: true,
		},
		{
			name: "unknown plugin type",
			builder: Builder{
				providerBuilders: map[string]providerBuilder{
					"test": {
						Import: "test",
					},
				},
			},
			pName: "unknown",
			config: map[string]interface{}{
				"test": "config",
			},
			expectedError: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := test.builder.BuildProvider(test.pName, test.config)
			if (err != nil) != test.expectedError {
				t.Errorf("BuildProvider() error = %v, wantErr %v", err, test.expectedError)
			}
		})
	}
}
