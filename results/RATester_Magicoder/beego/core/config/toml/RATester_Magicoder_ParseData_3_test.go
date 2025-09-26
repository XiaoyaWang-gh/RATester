package toml

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestParseData_3(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		c       *Config
		args    args
		want    config.Configer
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

			got, err := tt.c.ParseData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("Config.ParseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Config.ParseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
