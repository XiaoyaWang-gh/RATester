package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewBuilder_1(t *testing.T) {
	type args struct {
		client       *Client
		plugins      map[string]Descriptor
		localPlugins map[string]LocalDescriptor
	}
	tests := []struct {
		name    string
		args    args
		want    *Builder
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

			got, err := NewBuilder(tt.args.client, tt.args.plugins, tt.args.localPlugins)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewBuilder() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}
