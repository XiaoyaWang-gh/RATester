package memcache

import (
	"fmt"
	"testing"
)

func TestStartAndGC_3(t *testing.T) {
	type args struct {
		config string
	}
	tests := []struct {
		name    string
		rc      *Cache
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid json",
			rc:   &Cache{},
			args: args{
				config: `{"conn": "127.0.0.1:11211"}`,
			},
			wantErr: false,
		},
		{
			name: "Test with invalid json",
			rc:   &Cache{},
			args: args{
				config: `{"conn": 127.0.0.1:11211"}`,
			},
			wantErr: true,
		},
		{
			name: "Test with missing conn field",
			rc:   &Cache{},
			args: args{
				config: `{}`,
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

			rc := &Cache{}
			if err := rc.StartAndGC(tt.args.config); (err != nil) != tt.wantErr {
				t.Errorf("Cache.StartAndGC() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
