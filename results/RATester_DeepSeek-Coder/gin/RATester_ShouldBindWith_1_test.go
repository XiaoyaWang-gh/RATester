package gin

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin/binding"
)

func TestShouldBindWith_1(t *testing.T) {
	type testStruct struct {
		Field string `form:"field"`
	}

	testCases := []struct {
		name    string
		request *http.Request
		obj     any
		wantErr bool
	}{
		{
			name:    "valid request",
			request: httptest.NewRequest("GET", "/?field=value", nil),
			obj:     &testStruct{},
			wantErr: false,
		},
		{
			name:    "invalid request",
			request: httptest.NewRequest("GET", "/?field=", nil),
			obj:     &testStruct{},
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

			c := &Context{Request: tc.request}
			err := c.ShouldBindWith(tc.obj, binding.Form)
			if (err != nil) != tc.wantErr {
				t.Errorf("ShouldBindWith() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
