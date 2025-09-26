package gateway

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewExternalClusterClient_2(t *testing.T) {
	type args struct {
		endpoint   string
		caFilePath string
		token      types.FileOrContent
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				endpoint:   "https://example.com",
				caFilePath: "testdata/ca.pem",
				token:      types.FileOrContent("testdata/token"),
			},
			wantErr: false,
		},
		{
			name: "missing endpoint",
			args: args{
				endpoint:   "",
				caFilePath: "testdata/ca.pem",
				token:      types.FileOrContent("testdata/token"),
			},
			wantErr: true,
		},
		{
			name: "invalid token",
			args: args{
				endpoint:   "https://example.com",
				caFilePath: "testdata/ca.pem",
				token:      types.FileOrContent("testdata/non-existent-token"),
			},
			wantErr: true,
		},
		{
			name: "invalid ca file",
			args: args{
				endpoint:   "https://example.com",
				caFilePath: "testdata/non-existent-ca.pem",
				token:      types.FileOrContent("testdata/token"),
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

			_, err := newExternalClusterClient(tt.args.endpoint, tt.args.caFilePath, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("newExternalClusterClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
