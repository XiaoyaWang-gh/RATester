package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetSection_8(t *testing.T) {
	type args struct {
		section string
	}
	tests := []struct {
		name    string
		b       *beegoAppConfig
		args    args
		want    map[string]string
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

			got, err := tt.b.GetSection(tt.args.section)
			if (err != nil) != tt.wantErr {
				t.Errorf("beegoAppConfig.GetSection() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("beegoAppConfig.GetSection() = %v, want %v", got, tt.want)
			}
		})
	}
}
