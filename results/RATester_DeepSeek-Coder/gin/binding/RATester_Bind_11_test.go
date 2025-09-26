package binding

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBind_11(t *testing.T) {
	type testStruct struct {
		Field string `json:"field"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "valid yaml",
			body:    `field: value`,
			wantErr: false,
		},
		{
			name:    "invalid yaml",
			body:    `field: value`,
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

			req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tc.body))
			req.Header.Set("Content-Type", "application/yaml")

			obj := &testStruct{}
			err := yamlBinding{}.Bind(req, obj)

			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
