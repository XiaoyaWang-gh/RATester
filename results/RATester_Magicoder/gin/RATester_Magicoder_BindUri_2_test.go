package gin

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBindUri_2(t *testing.T) {
	type testStruct struct {
		Field1 string `uri:"field1"`
		Field2 int    `uri:"field2"`
	}

	testCases := []struct {
		name    string
		uri     string
		wantErr bool
	}{
		{
			name:    "valid uri",
			uri:     "/test/123/456",
			wantErr: false,
		},
		{
			name:    "invalid uri",
			uri:     "/test/abc/def",
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

			ctx := &Context{
				Request: &http.Request{
					URL: &url.URL{
						Path: tc.uri,
					},
				},
			}

			obj := &testStruct{}
			err := ctx.BindUri(obj)

			if tc.wantErr {
				if err == nil {
					t.Errorf("Expected error, but got nil")
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
			}
		})
	}
}
