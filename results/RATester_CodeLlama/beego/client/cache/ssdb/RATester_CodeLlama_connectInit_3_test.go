package ssdb

import (
	"fmt"
	"testing"
)

func TestConnectInit_3(t *testing.T) {
	type args struct {
		rc *Cache
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test connectInit",
			args: args{
				rc: &Cache{
					conninfo: []string{"127.0.0.1:8888"},
				},
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

			if err := tt.args.rc.connectInit(); (err != nil) != tt.wantErr {
				t.Errorf("connectInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
