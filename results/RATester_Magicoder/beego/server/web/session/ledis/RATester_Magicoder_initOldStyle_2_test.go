package ledis

import (
	"fmt"
	"testing"
)

func TestinitOldStyle_2(t *testing.T) {
	type args struct {
		cfgStr string
	}
	tests := []struct {
		name    string
		lp      *Provider
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			lp:   &Provider{},
			args: args{
				cfgStr: "test_path,1",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			lp:   &Provider{},
			args: args{
				cfgStr: "test_path",
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			lp:   &Provider{},
			args: args{
				cfgStr: "test_path,invalid_db",
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

			if err := tt.lp.initOldStyle(tt.args.cfgStr); (err != nil) != tt.wantErr {
				t.Errorf("Provider.initOldStyle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
