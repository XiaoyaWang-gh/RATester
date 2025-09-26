package binding

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBind_3(t *testing.T) {
	b := plainBinding{}

	testCases := []struct {
		name    string
		req     *http.Request
		obj     interface{}
		wantErr bool
	}{
		{
			name: "success",
			req: &http.Request{
				Body: ioutil.NopCloser(strings.NewReader(`{"field_name": "value"}`)),
			},
			obj:     &struct{ FieldName string }{},
			wantErr: false,
		},
		{
			name: "error",
			req: &http.Request{
				Body: ioutil.NopCloser(strings.NewReader(`{"field_name": "value"}`)),
			},
			obj:     &struct{ FieldName string }{},
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

			err := b.Bind(tc.req, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
				return
			}
		})
	}
}
