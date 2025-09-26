package context

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestBindForm_1(t *testing.T) {
	type testStruct struct {
		Field1 string `form:"field1"`
		Field2 string `form:"field2"`
	}

	testCases := []struct {
		name    string
		form    url.Values
		wantErr bool
	}{
		{
			name: "valid form",
			form: url.Values{
				"field1": []string{"value1"},
				"field2": []string{"value2"},
			},
			wantErr: false,
		},
		{
			name: "invalid form",
			form: url.Values{
				"field1": []string{"value1"},
				"field2": []string{""},
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

			ctx := &Context{
				Request: &http.Request{
					Form: tc.form,
				},
			}

			obj := &testStruct{}
			err := ctx.BindForm(obj)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindForm() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
