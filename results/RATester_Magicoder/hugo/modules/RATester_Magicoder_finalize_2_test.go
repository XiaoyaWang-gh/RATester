package modules

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
)

func Testfinalize_2(t *testing.T) {
	type args struct {
		logger loggers.Logger
	}
	tests := []struct {
		name    string
		m       *ModulesConfig
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

			if err := tt.m.finalize(tt.args.logger); (err != nil) != tt.wantErr {
				t.Errorf("ModulesConfig.finalize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
