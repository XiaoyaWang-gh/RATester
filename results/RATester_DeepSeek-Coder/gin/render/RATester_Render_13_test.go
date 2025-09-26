package render

import (
	"fmt"
	"net/http/httptest"
	"testing"
)

func TestRender_13(t *testing.T) {
	testCases := []struct {
		name    string
		data    Data
		wantErr bool
	}{
		{
			name: "success",
			data: Data{
				ContentType: "application/json",
				Data:        []byte(`{"key": "value"}`),
			},
			wantErr: false,
		},
		{
			name: "failure",
			data: Data{
				ContentType: "application/json",
				Data:        []byte(`invalid json`),
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

			w := httptest.NewRecorder()
			err := tc.data.Render(w)
			if (err != nil) != tc.wantErr {
				t.Errorf("Render() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
