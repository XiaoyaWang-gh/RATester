package ssdb

import (
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestConnectInit_5(t *testing.T) {
	type fields struct {
		client      *ssdb.Client
		Host        string
		Port        int
		maxLifetime int64
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				Host: "",
				Port: 0,
			},
			wantErr: true,
		},
		{
			name: "Test Case 2",
			fields: fields{
				Host: "localhost",
				Port: 8888,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			p := &Provider{
				client:      tt.fields.client,
				Host:        tt.fields.Host,
				Port:        tt.fields.Port,
				maxLifetime: tt.fields.maxLifetime,
			}
			if err := p.connectInit(); (err != nil) != tt.wantErr {
				t.Errorf("Provider.connectInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
