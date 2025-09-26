package memcache

import (
	"fmt"
	"testing"
)

func TestConnectInit_6(t *testing.T) {
	type fields struct {
		maxlifetime int64
		conninfo    []string
		poolsize    int
		password    string
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

			rp := &MemProvider{
				maxlifetime: tt.fields.maxlifetime,
				conninfo:    tt.fields.conninfo,
				poolsize:    tt.fields.poolsize,
				password:    tt.fields.password,
			}
			if err := rp.connectInit(); (err != nil) != tt.wantErr {
				t.Errorf("MemProvider.connectInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
