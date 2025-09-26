package binding

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBind_5(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 int    `form:"field2"`
	}

	testCases := []struct {
		name    string
		req     *http.Request
		obj     any
		wantErr bool
	}{
		{
			name: "valid form",
			req: &http.Request{
				Form: url.Values{
					"field1": []string{"value1"},
					"field2": []string{"123"},
				},
			},
			obj: &testStruct{},
		},
		{
			name: "invalid form",
			req: &http.Request{
				Form: url.Values{
					"field1": []string{"value1"},
					"field2": []string{"not_an_int"},
				},
			},
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

			formBinding := formBinding{}
			err := formBinding.Bind(tc.req, tc.obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("Bind() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
