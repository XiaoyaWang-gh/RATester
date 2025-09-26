package runtime

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestAddError_1(t *testing.T) {
	type testCase struct {
		name     string
		service  *UDPServiceInfo
		err      error
		critical bool
		want     *UDPServiceInfo
	}

	testCases := []testCase{
		{
			name: "add new error",
			service: &UDPServiceInfo{
				Err:    []string{"old error"},
				Status: StatusWarning,
			},
			err:      errors.New("new error"),
			critical: false,
			want: &UDPServiceInfo{
				Err:    []string{"old error", "new error"},
				Status: StatusWarning,
			},
		},
		{
			name: "add existing error",
			service: &UDPServiceInfo{
				Err:    []string{"new error"},
				Status: StatusWarning,
			},
			err:      errors.New("new error"),
			critical: false,
			want: &UDPServiceInfo{
				Err:    []string{"new error"},
				Status: StatusWarning,
			},
		},
		{
			name: "add critical error",
			service: &UDPServiceInfo{
				Err:    []string{"old error"},
				Status: StatusWarning,
			},
			err:      errors.New("critical error"),
			critical: true,
			want: &UDPServiceInfo{
				Err:    []string{"old error", "critical error"},
				Status: StatusDisabled,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tc.service.AddError(tc.err, tc.critical)
			if !reflect.DeepEqual(tc.service, tc.want) {
				t.Errorf("got %v, want %v", tc.service, tc.want)
			}
		})
	}
}
