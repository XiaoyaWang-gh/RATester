package alils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListConfig_1(t *testing.T) {
	type args struct {
		offset int
		size   int
	}
	tests := []struct {
		name         string
		p            *LogProject
		args         args
		wantCfgNames []string
		wantTotal    int
		wantErr      bool
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

			gotCfgNames, gotTotal, err := tt.p.ListConfig(tt.args.offset, tt.args.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogProject.ListConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfgNames, tt.wantCfgNames) {
				t.Errorf("LogProject.ListConfig() gotCfgNames = %v, want %v", gotCfgNames, tt.wantCfgNames)
			}
			if gotTotal != tt.wantTotal {
				t.Errorf("LogProject.ListConfig() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}
