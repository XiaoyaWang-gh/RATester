package hugolib

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/common/loggers"
	"github.com/gohugoio/hugo/config"
)

func TestsetMetaPre_1(t *testing.T) {
	type args struct {
		pi     *contentParseInfo
		logger loggers.Logger
		conf   config.AllProvider
	}
	tests := []struct {
		name    string
		p       *pageMeta
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

			if err := tt.p.setMetaPre(tt.args.pi, tt.args.logger, tt.args.conf); (err != nil) != tt.wantErr {
				t.Errorf("setMetaPre() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
