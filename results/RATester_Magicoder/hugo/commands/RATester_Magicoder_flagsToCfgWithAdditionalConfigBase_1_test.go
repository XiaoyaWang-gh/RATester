package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/simplecobra"
	"github.com/gohugoio/hugo/config"
)

func TestflagsToCfgWithAdditionalConfigBase_1(t *testing.T) {
	type args struct {
		cd                   *simplecobra.Commandeer
		cfg                  config.Provider
		additionalConfigBase string
	}
	tests := []struct {
		name string
		args args
		want config.Provider
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

			if got := flagsToCfgWithAdditionalConfigBase(tt.args.cd, tt.args.cfg, tt.args.additionalConfigBase); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flagsToCfgWithAdditionalConfigBase() = %v, want %v", got, tt.want)
			}
		})
	}
}
