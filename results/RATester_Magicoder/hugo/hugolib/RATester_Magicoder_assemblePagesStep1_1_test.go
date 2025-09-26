package hugolib

import (
	"context"
	"fmt"
	"testing"
)

func TestassemblePagesStep1_1(t *testing.T) {
	type fields struct {
		Site *Site
	}
	type args struct {
		ctx context.Context
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

			sa := &sitePagesAssembler{
				Site: tt.fields.Site,
			}
			if err := sa.assemblePagesStep1(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("sitePagesAssembler.assemblePagesStep1() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
