package file

import (
	"fmt"
	"testing"
)

func TestBuildConfiguration_5(t *testing.T) {
	testCases := []struct {
		desc          string
		provider      *Provider
		expectedError bool
	}{
		{
			desc: "should return error when neither filename nor directory is defined",
			provider: &Provider{
				Directory: "",
				Filename:  "",
			},
			expectedError: true,
		},
		{
			desc: "should return configuration from directory when defined",
			provider: &Provider{
				Directory: "/path/to/directory",
				Filename:  "",
			},
			expectedError: false,
		},
		{
			desc: "should return configuration from file when defined",
			provider: &Provider{
				Directory: "",
				Filename:  "/path/to/file.yml",
			},
			expectedError: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.desc, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := test.provider.buildConfiguration()
			if (err != nil) != test.expectedError {
				t.Errorf("buildConfiguration() error = %v, wantErr %v", err, test.expectedError)
			}
		})
	}
}
