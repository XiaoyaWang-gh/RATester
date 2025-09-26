package ssdb

import (
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestConnectInit_3(t *testing.T) {
	type fields struct {
		conn     *ssdb.Client
		conninfo []string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "Test Case 1",
			fields: fields{
				conn:     nil,
				conninfo: []string{"localhost:8888"},
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			fields: fields{
				conn:     nil,
				conninfo: []string{"localhost:abc"},
			},
			wantErr: true,
		},
		{
			name: "Test Case 3",
			fields: fields{
				conn:     nil,
				conninfo: []string{"localhost"},
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

			rc := &Cache{
				conn:     tt.fields.conn,
				conninfo: tt.fields.conninfo,
			}
			if err := rc.connectInit(); (err != nil) != tt.wantErr {
				t.Errorf("connectInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
