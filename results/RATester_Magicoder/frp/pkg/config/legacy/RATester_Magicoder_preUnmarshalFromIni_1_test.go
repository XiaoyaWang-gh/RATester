package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestPreUnmarshalFromIni_1(t *testing.T) {
	type args struct {
		cfg     ProxyConf
		prefix  string
		name    string
		section *ini.Section
	}
	tests := []struct {
		name    string
		args    args
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

			if err := preUnmarshalFromIni(tt.args.cfg, tt.args.prefix, tt.args.name, tt.args.section); (err != nil) != tt.wantErr {
				t.Errorf("preUnmarshalFromIni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
