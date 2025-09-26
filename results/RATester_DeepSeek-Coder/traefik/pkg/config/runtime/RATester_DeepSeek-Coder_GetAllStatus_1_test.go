package runtime

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetAllStatus_1(t *testing.T) {
	tests := []struct {
		name           string
		serverStatus   map[string]string
		expectedStatus map[string]string
	}{
		{
			name: "no servers",
			serverStatus: map[string]string{
				"server1": "status1",
				"server2": "status2",
			},
			expectedStatus: nil,
		},
		{
			name: "with servers",
			serverStatus: map[string]string{
				"server1": "status1",
				"server2": "status2",
			},
			expectedStatus: map[string]string{
				"server1": "status1",
				"server2": "status2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &ServiceInfo{
				serverStatus: test.serverStatus,
			}

			status := s.GetAllStatus()

			if !reflect.DeepEqual(status, test.expectedStatus) {
				t.Errorf("expected %v, got %v", test.expectedStatus, status)
			}
		})
	}
}
