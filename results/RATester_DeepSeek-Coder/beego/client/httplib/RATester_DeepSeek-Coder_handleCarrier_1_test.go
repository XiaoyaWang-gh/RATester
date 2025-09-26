package httplib

import (
	"fmt"
	"testing"
)

func TestHandleCarrier_1(t *testing.T) {
	type args struct {
		value interface{}
		req   *BeegoHTTPRequest
	}
	tests := []struct {
		name    string
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

			c := &Client{}
			if err := c.handleCarrier(tt.args.value, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("handleCarrier() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
