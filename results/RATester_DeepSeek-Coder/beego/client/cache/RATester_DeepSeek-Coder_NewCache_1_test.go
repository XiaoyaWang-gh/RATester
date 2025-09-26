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
			name: "TestNewCache_Success",
			args: args{
				adapterName: "file",
				config:      "{....}",
			},
			wantErr: false,
		},
		{
			name: "TestNewCache_UnknownAdapter",
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

			got, err := NewCache(tt.args.adapterName, tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCache() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (got != nil) != tt.wantErr {
				t.Errorf("NewCache() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}
