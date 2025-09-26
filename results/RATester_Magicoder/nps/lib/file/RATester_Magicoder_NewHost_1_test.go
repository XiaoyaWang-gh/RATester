package file

import (
	"fmt"
	"testing"
)

func TestNewHost_1(t *testing.T) {
	db := &DbUtils{
		JsonDb: &JsonDb{},
	}

	tests := []struct {
		name    string
		host    *Host
		wantErr bool
	}{
		{
			name: "TestNewHost_Success",
			host: &Host{
				Id:     1,
				Host:   "test.com",
				Scheme: "http",
			},
			wantErr: false,
		},
		{
			name: "TestNewHost_Fail_Host_Exist",
			host: &Host{
				Id:     1,
				Host:   "test.com",
				Scheme: "http",
			},
			wantErr: true,
		},
		{
			name: "TestNewHost_Fail_Location_Empty",
			host: &Host{
				Id:     2,
				Host:   "test.com",
				Scheme: "http",
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

			if err := db.NewHost(tt.host); (err != nil) != tt.wantErr {
				t.Errorf("NewHost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
