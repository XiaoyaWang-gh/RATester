package render

import (
	"errors"
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_13(t *testing.T) {
	testCases := []struct {
		name        string
		data        Data
		expectedErr error
	}{
		{
			name: "Success",
			data: Data{
				ContentType: "application/json",
				Data:        []byte(`{"key": "value"}`),
			},
			expectedErr: nil,
		},
		{
			name: "Error",
			data: Data{
				ContentType: "application/json",
				Data:        []byte(`{"key": "value"}`),
			},
			expectedErr: errors.New("some error"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			w := httptest.NewRecorder()
			err := tc.data.Render(w)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
			}
		})
	}
}
