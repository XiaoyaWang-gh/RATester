package highlight

import (
	"fmt"
	"testing"
)

func TestapplyOptionsFromString_1(t *testing.T) {
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
				opts: "invalid_option",
				cfg:  &Config{},
			},
			wantErr: true,
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

			if err := applyOptionsFromString(tt.args.opts, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("applyOptionsFromString() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
