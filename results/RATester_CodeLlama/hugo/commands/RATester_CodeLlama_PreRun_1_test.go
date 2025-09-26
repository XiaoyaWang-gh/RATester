package commands

import (
	"fmt"
	"testing"

	"github.com/bep/simplecobra"
)

func TestPreRun_1(t *testing.T) {
	type fields struct {
		r *rootCommand
	}
	type args struct {
		cd  *simplecobra.Commandeer
		run *simplecobra.Commandeer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test",
			fields: fields{
				r: &rootCommand{},
			},
			args: args{
				cd:  &simplecobra.Commandeer{},
				run: &simplecobra.Commandeer{},
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

			c := &serverCommand{
				r: tt.fields.r,
			}
			if err := c.PreRun(tt.args.cd, tt.args.run); (err != nil) != tt.wantErr {
				t.Errorf("PreRun() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
