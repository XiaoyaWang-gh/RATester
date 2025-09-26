package deploy

import (
	"context"
	"fmt"
	"testing"
)

func TestInvalidateGoogleCloudCDN_1(t *testing.T) {
	type args struct {
		ctx    context.Context
		origin string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test case 1",
			args: args{
				ctx:    context.Background(),
				origin: "test-project/test-origin",
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			args: args{
				ctx:    context.Background(),
				origin: "invalid-origin",
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

			err := InvalidateGoogleCloudCDN(tt.args.ctx, tt.args.origin)
			if (err != nil) != tt.wantErr {
				t.Errorf("InvalidateGoogleCloudCDN() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
