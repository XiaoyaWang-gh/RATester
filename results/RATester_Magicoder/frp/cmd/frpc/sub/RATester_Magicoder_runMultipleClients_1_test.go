package sub

import (
	"fmt"
	"testing"
)

func TestRunMultipleClients_1(t *testing.T) {
	type args struct {
		cfgDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				cfgDir: "/path/to/config/dir",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				cfgDir: "/path/to/invalid/dir",
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

			if err := runMultipleClients(tt.args.cfgDir); (err != nil) != tt.wantErr {
				t.Errorf("runMultipleClients() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
