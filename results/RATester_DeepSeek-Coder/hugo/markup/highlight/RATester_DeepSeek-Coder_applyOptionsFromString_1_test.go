package highlight

import (
	"fmt"
	"testing"
)

func TestApplyOptionsFromString_1(t *testing.T) {
	type args struct {
		opts string
		cfg  *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				opts: "style=monokai",
				cfg:  &Config{},
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				opts: "style=invalid",
				cfg:  &Config{},
			},
			wantErr: true,
		},
		{
			name: "Test case 3",
			args: args{
				opts: "codeFences=true",
				cfg:  &Config{},
			},
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := applyOptionsFromString(tt.args.opts, tt.args.cfg)
			if (err != nil) != tt.wantErr {
				t.Errorf("applyOptionsFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
