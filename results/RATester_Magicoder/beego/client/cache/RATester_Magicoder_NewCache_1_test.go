package cache

import (
	"fmt"
	"testing"
)

func TestNewCache_1(t *testing.T) {
	type args struct {
		adapterName string
		config      string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				adapterName: "file",
				config:      "{....}",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				adapterName: "unknown",
				config:      "{....}",
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

			_, err := NewCache(tt.args.adapterName, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
