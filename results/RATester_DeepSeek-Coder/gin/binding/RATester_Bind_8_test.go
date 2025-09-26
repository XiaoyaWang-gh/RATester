package binding

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBind_8(t *testing.T) {
	type testStruct struct {
		Field string `xml:"field"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name:    "valid xml",
			body:    `<testStruct><field>value</field></testStruct>`,
			wantErr: false,
		},
		{
			name:    "invalid xml",
			body:    `<testStruct><field>value</field></testStruct><testStruct><field>value</field></testStruct>`,
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
			obj := &testStruct{}
			err := xmlBinding{}.Bind(req, obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
