package yaml

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestParse_1(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantY   config.Configer
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

			c := &Config{}
			gotY, err := c.Parse(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotY, tt.wantY) {
				t.Errorf("Parse() gotY = %v, want %v", gotY, tt.wantY)
			}
		})
	}
}
