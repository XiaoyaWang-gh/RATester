package deploy

import (
	"context"
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/deploy/deployconfig"
)

func TestOpenBucket_1(t *testing.T) {
	t.Parallel()

	ctx := context.Background()

	tests := []struct {
		name    string
		target  *deployconfig.Target
		wantErr bool
	}{
		{
			name: "success",
			target: &deployconfig.Target{
				Name: "test",
				URL:  "mem://test",
			},
			wantErr: false,
		},
		{
			name: "failure",
			target: &deployconfig.Target{
				Name: "test",
				URL:  "invalid://test",
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

			d := &Deployer{
				target: tt.target,
			}
			_, err := d.openBucket(ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("openBucket() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
