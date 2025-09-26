package collector

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestCreateBody_1(t *testing.T) {
	type args struct {
		staticConfiguration *static.Configuration
	}
	tests := []struct {
		name    string
		args    args
		want    *bytes.Buffer
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

			got, err := createBody(tt.args.staticConfiguration)
			if (err != nil) != tt.wantErr {
				t.Errorf("createBody() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createBody() = %v, want %v", got, tt.want)
			}
		})
	}
}
