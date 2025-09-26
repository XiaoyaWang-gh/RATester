package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewProvider_1(t *testing.T) {
	type args struct {
		builder      providerBuilder
		config       map[string]interface{}
		providerName string
	}
	tests := []struct {
		name    string
		args    args
		want    *Provider
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

			got, err := newProvider(tt.args.builder, tt.args.config, tt.args.providerName)
			if (err != nil) != tt.wantErr {
				t.Errorf("newProvider() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newProvider() = %v, want %v", got, tt.want)
			}
		})
	}
}
