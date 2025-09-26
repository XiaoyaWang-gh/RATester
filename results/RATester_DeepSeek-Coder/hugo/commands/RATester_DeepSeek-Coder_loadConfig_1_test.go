package commands

import (
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestLoadConfig_1(t *testing.T) {
	type args struct {
		cd      *simplecobra.Commandeer
		running bool
	}

	tests := []struct {
		name    string
		c       *hugoBuilder
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

			if err := tt.c.loadConfig(tt.args.cd, tt.args.running); (err != nil) != tt.wantErr {
				t.Errorf("hugoBuilder.loadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
