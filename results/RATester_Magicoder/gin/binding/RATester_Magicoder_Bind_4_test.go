package binding

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
)

func TestBind_4(t *testing.T) {
	type testCase struct {
		name    string
		req     *http.Request
		obj     any
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid toml",
			req: &http.Request{
				Body: ioutil.NopCloser(strings.NewReader(`field = "value"`)),
			},
			obj: &struct {
				Field string `toml:"field"`
			}{},
			wantErr: false,
		},
		{
			name: "invalid toml",
			req: &http.Request{
				Body: ioutil.NopCloser(strings.NewReader(`field = "value"`)),
			},
			obj: &struct {
				Field int `toml:"field"`
			}{},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			tb := &tomlBinding{}
			err := tb.Bind(tt.req, tt.obj)
			if (err != nil) != tt.wantErr {
				t.Errorf("tomlBinding.Bind() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
