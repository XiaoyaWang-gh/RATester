package api

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/static"
)

func TestGetFeatures_1(t *testing.T) {
	type args struct {
		conf static.Configuration
	}
	tests := []struct {
		name string
		args args
		want features
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

			if got := getFeatures(tt.args.conf); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFeatures() = %v, want %v", got, tt.want)
			}
		})
	}
}
