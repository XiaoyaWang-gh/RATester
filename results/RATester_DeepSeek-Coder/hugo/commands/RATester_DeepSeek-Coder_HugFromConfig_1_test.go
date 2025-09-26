package commands

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/hugolib"
)

func TestHugFromConfig_1(t *testing.T) {
	type args struct {
		conf *commonConfig
	}
	tests := []struct {
		name    string
		r       *rootCommand
		args    args
		want    *hugolib.HugoSites
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

			got, err := tt.r.HugFromConfig(tt.args.conf)
			if (err != nil) != tt.wantErr {
				t.Errorf("rootCommand.HugFromConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rootCommand.HugFromConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
