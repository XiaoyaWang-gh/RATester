package ingress

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/types"
)

func TestNewExternalClusterClient_1(t *testing.T) {
	type args struct {
		endpoint   string
		caFilePath string
		token      types.FileOrContent
	}
	tests := []struct {
		name    string
		args    args
		want    *clientWrapper
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				endpoint:   "https://example.com",
				caFilePath: "/path/to/ca.crt",
				token:      "token",
			},
			want: &clientWrapper{
				// Fill in the fields you want to test
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				endpoint:   "",
				caFilePath: "/path/to/ca.crt",
				token:      "token",
			},
			want:    nil,
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

			got, err := newExternalClusterClient(tt.args.endpoint, tt.args.caFilePath, tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("newExternalClusterClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newExternalClusterClient() = %v, want %v", got, tt.want)
			}
		})
	}
}
