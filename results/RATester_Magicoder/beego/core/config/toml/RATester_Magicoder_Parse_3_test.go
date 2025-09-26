package toml

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestParse_3(t *testing.T) {
	type args struct {
		filename string
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

			got, err := tt.c.Parse(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
