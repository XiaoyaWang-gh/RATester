package plugins

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateConfig_1(t *testing.T) {
	type args struct {
		config map[string]interface{}
	}
	tests := []struct {
		name    string
		b       yaegiMiddlewareBuilder
		args    args
		want    reflect.Value
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

			got, err := tt.b.createConfig(tt.args.config)
			if (err != nil) != tt.wantErr {
				t.Errorf("createConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
