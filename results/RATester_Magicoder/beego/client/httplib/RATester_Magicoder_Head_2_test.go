package httplib

import (
	"fmt"
	"testing"
)

func TestHead_2(t *testing.T) {
	client := &Client{
		Name:     "test",
		Endpoint: "http://localhost:8080",
	}

	type testCase struct {
		name    string
		value   interface{}
		path    string
		opts    []BeegoHTTPRequestOption
		wantErr bool
	}

	tests := []testCase{
		{
			name:    "success",
			value:   nil,
			path:    "/test",
			opts:    []BeegoHTTPRequestOption{},
			wantErr: false,
		},
		{
			name:    "failure",
			value:   nil,
			path:    "/test",
			opts:    []BeegoHTTPRequestOption{},
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

			err := client.Head(tt.value, tt.path, tt.opts...)
			if (err != nil) != tt.wantErr {
				t.Errorf("Head() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
