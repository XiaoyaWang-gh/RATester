package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseFile_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		ini     *IniConfig
		args    args
		want    *IniConfigContainer
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

			got, err := tt.ini.parseFile(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("IniConfig.parseFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IniConfig.parseFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
