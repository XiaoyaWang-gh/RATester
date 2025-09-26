package commands

import (
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestLoadConfig_1(t *testing.T) {
	type fields struct {
		r *rootCommand
		s *serverCommand
	}
	type args struct {
		cd      *simplecobra.Commandeer
		running bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test_case_1",
			fields: fields{
				r: &rootCommand{},
				s: &serverCommand{},
			},
			args: args{
				cd:      &simplecobra.Commandeer{},
				running: true,
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

			c := &hugoBuilder{
				r: tt.fields.r,
				s: tt.fields.s,
			}
			if err := c.loadConfig(tt.args.cd, tt.args.running); (err != nil) != tt.wantErr {
				t.Errorf("loadConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
