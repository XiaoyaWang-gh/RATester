package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestUnmarshalFromIni_1(t *testing.T) {
	type args struct {
		prefix  string
		name    string
		section *ini.Section
	}
	tests := []struct {
		name    string
		cfg     *XTCPVisitorConf
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

			err := tt.cfg.UnmarshalFromIni(tt.args.prefix, tt.args.name, tt.args.section)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalFromIni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
