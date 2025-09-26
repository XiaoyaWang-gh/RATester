package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseData_2(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Configer
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

			ini := &IniConfig{}
			got, err := ini.ParseData(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
