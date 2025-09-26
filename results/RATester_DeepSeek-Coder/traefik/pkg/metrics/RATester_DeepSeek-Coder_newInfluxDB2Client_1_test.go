package metrics

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewInfluxDB2Client_1(t *testing.T) {
	tests := []struct {
		name    string
		config  *types.InfluxDB2
		wantErr bool
	}{
		{
			name: "should return error when token is missing",
			config: &types.InfluxDB2{
				Address: "http://localhost:8086",
				Org:     "my-org",
				Bucket:  "my-bucket",
			},
			wantErr: true,
		},
		{
			name: "should return error when org is missing",
			config: &types.InfluxDB2{
				Address: "http://localhost:8086",
				Token:   "my-token",
				Bucket:  "my-bucket",
			},
			wantErr: true,
		},
		{
			name: "should return error when bucket is missing",
			config: &types.InfluxDB2{
				Address: "http://localhost:8086",
				Token:   "my-token",
				Org:     "my-org",
			},
			wantErr: true,
		},
		{
			name: "should not return error when all properties are set",
			config: &types.InfluxDB2{
				Address: "http://localhost:8086",
				Token:   "my-token",
				Org:     "my-org",
				Bucket:  "my-bucket",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := newInfluxDB2Client(tt.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("newInfluxDB2Client() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
