package modules

import (
	"fmt"
	"testing"
)

func TestinitModules_1(t *testing.T) {
	type fields struct {
		Client *Client
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &collector{
				Client: tt.fields.Client,
			}
			if err := c.initModules(); (err != nil) != tt.wantErr {
				t.Errorf("collector.initModules() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
