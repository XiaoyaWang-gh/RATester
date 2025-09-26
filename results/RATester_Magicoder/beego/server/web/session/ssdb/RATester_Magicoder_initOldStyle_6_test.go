package ssdb

import (
	"fmt"
	"testing"
)

func TestinitOldStyle_6(t *testing.T) {
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
			name: "Test case 1",
			p:    &Provider{},
			args: args{
				savePath: "localhost:8080",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
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

			if err := tt.p.initOldStyle(tt.args.savePath); (err != nil) != tt.wantErr {
				t.Errorf("Provider.initOldStyle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
