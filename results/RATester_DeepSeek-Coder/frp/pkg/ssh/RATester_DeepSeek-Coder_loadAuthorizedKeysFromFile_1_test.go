package ssh

import (
	"fmt"
	"os"
	"testing"
)

func TestLoadAuthorizedKeysFromFile_1(t *testing.T) {
	testCases := []struct {
		name          string
		filePath      string
		expectedError error
	}{
		{
			name:          "File does not exist",
			filePath:      "/non/existent/path",
			expectedError: os.ErrNotExist,
		},
		{
			name:          "File exists",
			filePath:      "testdata/authorized_keys",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := loadAuthorizedKeysFromFile(tc.filePath)
			if err != tc.expectedError {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
