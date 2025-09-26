package ssdb

import (
	"fmt"
	"testing"
)

func TestInitOldStyle_6(t *testing.T) {
	type args struct {
		savePath string
	}
	tests := []struct {
		name    string
		p       *Provider
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid address",
			p:    &Provider{},
			args: args{
				savePath: "localhost:8080",
			},
			wantErr: false,
		},
		{
			name: "Test with invalid address",
			p:    &Provider{},
			args: args{
				savePath: "localhost:abc",
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

			err := tt.p.initOldStyle(tt.args.savePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Provider.initOldStyle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
