package binding

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestBind_4(t *testing.T) {
	type testStruct struct {
		Field1 string `toml:"field1"`
		Field2 int    `toml:"field2"`
	}

	testCases := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{
			name: "valid toml",
			body: `field1 = "value1"
field2 = 2`,
			wantErr: false,
		},
		{
			name: "invalid toml",
			body: `field1 = "value1"
field2 = "not an int"`,
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
			err := tomlBinding{}.Bind(req, obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
