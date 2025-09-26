package try

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHasBody_1(t *testing.T) {
	tests := []struct {
		name    string
		res     *http.Response
		wantErr bool
	}{
		{
			name: "has body",
			res: &http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString("hello")),
			},
			wantErr: false,
		},
		{
			name: "no body",
			res: &http.Response{
				Body: ioutil.NopCloser(bytes.NewBufferString("")),
			},
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

			if err := HasBody()(tt.res); (err != nil) != tt.wantErr {
				t.Errorf("HasBody() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
