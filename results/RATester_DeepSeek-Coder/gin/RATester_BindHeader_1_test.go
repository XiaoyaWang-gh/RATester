package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestBindHeader_1(t *testing.T) {
	type testStruct struct {
		HeaderField string `header:"HeaderField"`
	}

	tests := []struct {
		name    string
		header  http.Header
		wantErr bool
	}{
		{
			name: "Test with valid header",
			header: http.Header{
				"HeaderField": []string{"testValue"},
			},
			wantErr: false,
		},
		{
			name:    "Test with empty header",
			header:  http.Header{},
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

			c := &Context{
				Request: &http.Request{
					Header: tt.header,
				},
			}

			s := &testStruct{}
			err := c.BindHeader(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("BindHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
