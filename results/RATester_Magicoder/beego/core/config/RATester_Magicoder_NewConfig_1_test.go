package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewConfig_1(t *testing.T) {
	type args struct {
		adapterName string
		filename    string
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

			got, err := NewConfig(tt.args.adapterName, tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
