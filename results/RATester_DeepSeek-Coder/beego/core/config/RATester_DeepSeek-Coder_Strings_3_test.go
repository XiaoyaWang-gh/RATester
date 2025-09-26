package config

import (
	"fmt"
	"reflect"
	"testing"
)

func TestStrings_3(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		c       *BaseConfiger
		args    args
		want    []string
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

			got, err := tt.c.Strings(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("BaseConfiger.Strings() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseConfiger.Strings() = %v, want %v", got, tt.want)
			}
		})
	}
}
