package json

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/core/config"
)

func TestParseData_4(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		js      *JSONConfig
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

			got, err := tt.js.ParseData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("JSONConfig.ParseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JSONConfig.ParseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
