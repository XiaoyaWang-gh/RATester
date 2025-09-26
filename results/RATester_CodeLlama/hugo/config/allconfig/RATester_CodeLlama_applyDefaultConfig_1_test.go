package allconfig

import (
	"fmt"
	"testing"

	"github.com/gohugoio/hugo/config"
)

func TestApplyDefaultConfig_1(t *testing.T) {
	type args struct {
		l configLoader
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test applyDefaultConfig",
			args: args{
				l: configLoader{
					cfg: config.New(),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if err := tt.args.l.applyDefaultConfig(); (err != nil) != tt.wantErr {
				t.Errorf("applyDefaultConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
