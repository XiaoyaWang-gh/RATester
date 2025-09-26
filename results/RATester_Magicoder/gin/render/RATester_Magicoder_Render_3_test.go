package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_3(t *testing.T) {
	testCases := []struct {
		name    string
		secure  SecureJSON
		wantErr bool
	}{
		{
			name: "Test Case 1",
			secure: SecureJSON{
				Prefix: "prefix",
				Data:   []string{"data1", "data2"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			secure: SecureJSON{
				Prefix: "prefix",
				Data:   "data",
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			secure: SecureJSON{
				Prefix: "prefix",
				Data:   map[string]string{"key": "value"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 4",
			secure: SecureJSON{
				Prefix: "prefix",
				Data:   nil,
			},
			wantErr: false,
		},
		{
			name: "Test Case 5",
			secure: SecureJSON{
				Prefix: "prefix",
				Data:   make(chan int),
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := tc.secure.Render(httptest.NewRecorder())
			if (err != nil) != tc.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
