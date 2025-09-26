package commands

import (
	"errors"
	"fmt"
	"testing"
)

func TestCreateServerPorts_1(t *testing.T) {
	type testCase struct {
		name          string
		serverCommand serverCommand
		expectedErr   error
	}

	testCases := []testCase{
		{
			name: "Test case 1",
			serverCommand: serverCommand{
				serverAppend: true,
			},
			expectedErr: nil,
		},
		{
			name: "Test case 2",
			serverCommand: serverCommand{
				serverAppend: false,
			},
			expectedErr: errors.New("--appendPort=false not supported when in multihost mode"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.serverCommand.createServerPorts(nil)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
		})
	}
}
