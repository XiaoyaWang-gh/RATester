package allconfig

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/config"
	"github.com/gohugoio/hugo/modules"
)

func TestLoadConfigMain_1(t *testing.T) {
	type args struct {
		d ConfigSourceDescriptor
	}
	tests := []struct {
		name    string
		l       *configLoader
		args    args
		want    config.LoadConfigResult
		want1   modules.ModulesConfig
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

			got, got1, err := tt.l.loadConfigMain(tt.args.d)
			if (err != nil) != tt.wantErr {
				t.Errorf("configLoader.loadConfigMain() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("configLoader.loadConfigMain() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("configLoader.loadConfigMain() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
