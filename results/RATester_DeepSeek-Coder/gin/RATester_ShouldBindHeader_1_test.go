package gin

import (
	"fmt"
	"net/http"
	"testing"
)

func TestShouldBindHeader_1(t *testing.T) {
	type testStruct struct {
		HeaderField string `header:"HeaderField"`
	}

	tests := []struct {
		name    string
		header  string
		wantErr bool
	}{
		{
			name:    "valid header",
			header:  "headerValue",
			wantErr: false,
		},
		{
			name:    "invalid header",
			header:  "",
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
					Header: make(http.Header),
				},
			}
			c.Request.Header.Set("HeaderField", tt.header)

			s := &testStruct{}
			err := c.ShouldBindHeader(s)
			if (err != nil) != tt.wantErr {
				t.Errorf("ShouldBindHeader() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
