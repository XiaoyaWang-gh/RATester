package ledis

import (
	"fmt"
	"testing"
)

func TestInitOldStyle_2(t *testing.T) {
	type args struct {
		cfgStr string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				cfgStr: "test_path,1",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				cfgStr: "test_path",
			},
			wantErr: false,
		},
		{
			name: "Test case 3",
			args: args{
				cfgStr: "test_path,invalid_db",
			},
			wantErr: true,
		},
		{
			name: "Test case 4",
			args: args{
				cfgStr: "test_path,",
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

			lp := &Provider{}
			if err := lp.initOldStyle(tt.args.cfgStr); (err != nil) != tt.wantErr {
				t.Errorf("Provider.initOldStyle() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
