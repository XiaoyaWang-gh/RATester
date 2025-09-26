package legacy

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadAllProxyConfsFromIni_1(t *testing.T) {
	type args struct {
		prefix string
		source interface{}
		start  []string
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]ProxyConf
		want1   map[string]VisitorConf
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

			got, got1, err := LoadAllProxyConfsFromIni(tt.args.prefix, tt.args.source, tt.args.start)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAllProxyConfsFromIni() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAllProxyConfsFromIni() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadAllProxyConfsFromIni() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
