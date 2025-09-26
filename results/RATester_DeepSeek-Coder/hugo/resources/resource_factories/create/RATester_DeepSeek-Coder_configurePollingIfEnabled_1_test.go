package create

import (
	"fmt"
	"net/http"
	"testing"
)

func TestConfigurePollingIfEnabled_1(t *testing.T) {
	type args struct {
		uri        string
		optionsKey string
		getRes     func() (*http.Response, error)
	}
	tests := []struct {
		name string
		args args
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

			c := &Client{}
			c.configurePollingIfEnabled(tt.args.uri, tt.args.optionsKey, tt.args.getRes)
		})
	}
}
