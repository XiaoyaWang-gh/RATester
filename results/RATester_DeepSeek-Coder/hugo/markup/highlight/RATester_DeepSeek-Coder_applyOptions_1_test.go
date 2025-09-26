package highlight

import (
	"fmt"
	"testing"
)

func TestApplyOptions_1(t *testing.T) {
	type args struct {
		opts any
		cfg  *Config
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with nil opts",
			args: args{
				opts: nil,
				cfg:  &Config{},
			},
			wantErr: false,
		},
		{
			name: "Test with map opts",
			args: args{
				opts: map[string]any{
					"Style": "monokai",
				},
				cfg: &Config{},
			},
			wantErr: false,
		},
		{
			name: "Test with string opts",
			args: args{
				opts: "monokai",
				cfg:  &Config{},
			},
			wantErr: false,
		},
		{
			name: "Test with invalid opts",
			args: args{
				opts: 123,
				cfg:  &Config{},
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

			if err := applyOptions(tt.args.opts, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("applyOptions() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
