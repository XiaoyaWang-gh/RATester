package commands

import (
	"fmt"
	"testing"
)

func TestFullBuild_1(t *testing.T) {
	type fields struct {
		r *rootCommand
	}
	type args struct {
		noBuildLock bool
	}
	tests := []struct {
		name    string
		fields  fields
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

			c := &hugoBuilder{
				r: tt.fields.r,
			}
			if err := c.fullBuild(tt.args.noBuildLock); (err != nil) != tt.wantErr {
				t.Errorf("hugoBuilder.fullBuild() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
