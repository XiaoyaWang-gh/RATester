package legacy

import (
	"fmt"
	"testing"

	"gopkg.in/ini.v1"
)

func TestPreVisitorUnmarshalFromIni_1(t *testing.T) {
	type args struct {
		cfg     VisitorConf
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

			if err := preVisitorUnmarshalFromIni(tt.args.cfg, tt.args.prefix, tt.args.name, tt.args.section); (err != nil) != tt.wantErr {
				t.Errorf("preVisitorUnmarshalFromIni() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
