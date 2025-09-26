package traefik

import (
	"fmt"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestGetRedirectPort_1(t *testing.T) {
	type args struct {
		name string
		def  *static.Redirections
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			i := &Provider{}
			got, err := i.getRedirectPort(tt.args.name, tt.args.def)
			if (err != nil) != tt.wantErr {
				t.Errorf("getRedirectPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getRedirectPort() = %v, want %v", got, tt.want)
			}
		})
	}
}
