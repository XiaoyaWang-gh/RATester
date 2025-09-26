package config

import (
	"fmt"
	"testing"
)

func TestInitGlobalInstance_1(t *testing.T) {
	type args struct {
		name string
		cfg  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				name: "test",
				cfg:  "config",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				name: "",
				cfg:  "config",
			},
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				name: "test",
				cfg:  "",
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

			if err := InitGlobalInstance(tt.args.name, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("InitGlobalInstance() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
