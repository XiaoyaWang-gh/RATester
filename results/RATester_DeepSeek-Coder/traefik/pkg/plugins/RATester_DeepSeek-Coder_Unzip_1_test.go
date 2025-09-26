package plugins

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

func TestUnzip_1(t *testing.T) {
	type fields struct {
		HTTPClient *http.Client
		baseURL    *url.URL
		archives   string
		stateFile  string
		goPath     string
		sources    string
	}
	type args struct {
		pName    string
		pVersion string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Client{
				HTTPClient: tt.fields.HTTPClient,
				baseURL:    tt.fields.baseURL,
				archives:   tt.fields.archives,
				stateFile:  tt.fields.stateFile,
				goPath:     tt.fields.goPath,
				sources:    tt.fields.sources,
			}
			if err := c.Unzip(tt.args.pName, tt.args.pVersion); (err != nil) != tt.wantErr {
				t.Errorf("Client.Unzip() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
