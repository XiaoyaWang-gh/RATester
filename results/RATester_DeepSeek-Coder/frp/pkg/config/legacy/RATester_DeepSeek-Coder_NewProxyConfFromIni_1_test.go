package legacy

import (
	"fmt"
	"reflect"
	"testing"

	"gopkg.in/ini.v1"
)

func TestNewProxyConfFromIni_1(t *testing.T) {
	type args struct {
		prefix  string
		name    string
		section *ini.Section
	}
	tests := []struct {
		name    string
		args    args
		want    ProxyConf
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

			got, err := NewProxyConfFromIni(tt.args.prefix, tt.args.name, tt.args.section)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewProxyConfFromIni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProxyConfFromIni() = %v, want %v", got, tt.want)
			}
		})
	}
}
